package main

import (
	"fmt"
	"math/rand"
)

type PKE struct {
	prime_modulus int
	generator     int
}

/*
 *  Generate generator ^ x % prime_modulus
 */
func (this *PKE) PowerMod(gen, x int) int {
	y := gen
	for x > 1 {
		y = (y * gen) % this.prime_modulus
		x--
	}
	return y
}

func RandomInRange(min, max int) int {
	return min + rand.Intn(max-min)
}

type DFExchange struct {
	name          string
	private_key   int  // private key  that user uses
	my_public_key int  // my public key
	shared_secret int  // secret shared between two user
	pke           *PKE // public shared
}

func Construct(name string, pke *PKE) DFExchange {
	var d DFExchange
	d.name = name
	d.pke = pke
	d.private_key = RandomInRange(2, pke.prime_modulus-1)
	d.my_public_key = pke.PowerMod(pke.generator, d.private_key)
	return d
}

func (this *DFExchange) GetPublicKey() int {
	return this.my_public_key
}

// Generate the shared secret by PowerMod
func (this *DFExchange) GenerateSharedSecretKey(peerpk int) {
	fmt.Println(this.name, "gets public key", peerpk, "that is used to generate a shared secret key")
	this.shared_secret = this.pke.PowerMod(peerpk, this.private_key)
}

// simple XOR
func (this *DFExchange) Encrypt(x int) int {
	y := x ^ this.shared_secret
	fmt.Println(this.name, "encrypts data from '", x, "' to '", y, "'")
	return y
}

func (this *DFExchange) Decrypt(x int) int {
	y := x ^ this.shared_secret
	fmt.Println(this.name, "decrypts data from '", x, "' to '", y, "'")
	return y
}

func main() {

	pke := PKE{17, 3}
	alice := Construct("alice", &pke)
	bob := Construct("bob", &pke)
	// alice sends her public key to Bob and Bob sends his to alice
	fmt.Println("Alice Public Key:", alice.GetPublicKey())
	fmt.Println("Bob Public Key:", bob.GetPublicKey())
	bob.GenerateSharedSecretKey(alice.GetPublicKey())
	alice.GenerateSharedSecretKey(bob.GetPublicKey())
	// alice encrypts her data that bob decrypts
	bob.Decrypt(alice.Encrypt(10))

}
