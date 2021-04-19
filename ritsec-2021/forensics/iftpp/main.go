package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"iftpp/iftpp"
	"log"
	"os"
	"sort"

	"github.com/golang/protobuf/proto"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var plaintextChecksum []byte

// calculate the shared key by combining keys, sort by
// descending, then taking sha1
func calcSharedKey(key1 []byte, key2 []byte) []byte {
	combined := append(key1, key2...) // put two keys together
	sort.Slice(combined, func(i int, j int) bool {
		return combined[i] > combined[j]
	}) // sort descending
	hasher := sha1.New()
	hasher.Write(combined)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return []byte(sha)
}

func verifyChecksum(pb *iftpp.IFTPP, plainData []byte) bool {
	if pb != nil {
		h := sha1.New()
		h.Write(pb.B_1)
		hash := h.Sum(nil)
		//fmt.Printf("%x\n", hash)

		// sha := sha1.Sum(pb.B_1)
		// sha = sha1.
		b64Data := base64.StdEncoding.EncodeToString(hash)
		//fmt.Println(b64Data)

		// Delete last char of the b64 string
		b64Data = b64Data[:len(b64Data)-1]
		// Last 8 chars are the checksum
		checksum := b64Data[len(b64Data)-8:]
		//fmt.Println(checksum)
		//fmt.Println(string(pb.B_2))
		str_sum := string(pb.B_2)
		if checksum != str_sum {
			// Checksums match which is a big win
			return false
		}
		// Checksums dont match so fuck me am I right?
		return true

	} else {
		// Checksum for total file
		h := sha1.New()
		h.Write(plainData)
		hash := h.Sum(nil)
		b64Data := base64.StdEncoding.EncodeToString(hash)
		//fmt.Println(b64Data)

		// Delete last char of the b64 string
		b64Data = b64Data[:len(b64Data)-1]
		// Last 8 chars are the checksum
		plainchecksum := b64Data[len(b64Data)-8:]
		//fmt.Println(checksum)
		//fmt.Println(string(pb.B_2))
		str_sum := string(plainchecksum)
		str_fullFileChecksum := string(plaintextChecksum)
		if str_fullFileChecksum != str_sum {
			// Checksums match which is a big win
			return false
		}
		// Checksums dont match so fuck me am I right?
		return true
	}
}

func main() {
	fmt.Println("Hello World")
	handle, err := pcap.OpenOffline("../iftpp_challenge.pcap")
	if err != nil {
		panic(err)
	}

	var data [][]uint8
	var jpeg []uint8
	var key1 []byte
	var key2 []byte

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//handlePacket(packet)  // Do something with a packet here.
		pb := &iftpp.IFTPP{}
		icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
		if icmpLayer == nil {
			continue
		}

		buf := icmpLayer.LayerPayload()[2:len(icmpLayer.LayerPayload())]

		if err = proto.Unmarshal(buf, pb); err != nil {
			log.Println("Failed to parse ICMP data to protobuf: " + err.Error())
		}

		/*if pb.TypeFlag != 1 {
			// Validate Checksum
			b_err := verifyChecksum(pb, nil)
			if !b_err {
				fmt.Println("Rat bastard")
				continue
			}
		} else {
			continue
		}*/

		switch pb.TypeFlag {
		case 2:
			key1 = pb.B_1
			//key1 = append(pb.B_1)//, pb.B_2...)
		case 3:
			key2 = pb.B_1
			//key2 = append(key2, pb.B_2...)
			//key2 = append(key2, pb.B_2...)
		case 5:
			data = append(data, pb.B_1)
		case 6:
			fmt.Println("Getting full checksum")
			plaintextChecksum = pb.B_2
		}
	}

	log.Println("Now for the work")
	shared_key := calcSharedKey(key1, key2)
	log.Println(shared_key)

	// Decrypt Data
	for i := 0; i < len(data); i++ {
		var decryptedData []byte
		// Data[i][] --> packet
		// Data[i][y] --> byte array for packet x
		for x := 0; x < len(data[i]); x++ {
			key_c := shared_key[x%(len(shared_key))]
			xor_out := data[i][x] ^ key_c
			// xor_out := data[i][x] ^ shared_key[x % len(shared_key)]
			decryptedData = append(decryptedData, xor_out)
		}

		jpeg = append(jpeg, decryptedData...)
	}

	// Verify Plaintext with Checksum from FLAG TYPE 6
	//b_err := verifyChecksum(nil, jpeg)
	//if !b_err {
	//	log.Fatalln("Plaintext checksum failed...")
	//}

	f, err := os.Create("./iftpp_flag.jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(jpeg)
	if err != nil {
		panic(err)
	}

	f.Sync()

}
