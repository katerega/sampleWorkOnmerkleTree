package main

import (
  "crypto/sha256"
  "log"

  "github.com/katerega/merkletree"
)

//implement interface and content.
type TestContent struct {
  x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
  h := sha256.New()
  if _, err := h.Write([]byte(t.x)); err != nil {
    return nil, err
  }

  return h.Sum(nil), nil
}

//Equals tests two contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
  return t.x == other.(TestContent).x, nil
}

func main() {
  //Build list of Content to build tree
  var list []merkletree.Content
  list = append(list, TestContent{x: "Gardencordinates"})
  list = append(list, TestContent{x: "Croptype"})
  list = append(list, TestContent{x: "durationToharvest"})
  list = append(list, TestContent{x: "Contributionamount"})

  //Create a new Merkle Tree from the list of Content
  t, err := merkletree.NewTree(list)
  if err != nil {
    log.Fatal(err)
  }

  //Get Root of the tree
  mr := t.MerkleRoot()
  log.Println(mr)

  //has the node
  vt, err := t.VerifyTree()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Verify Tree: ", vt)

  //Verify content in the tree
  vc, err := t.VerifyContent(list[0])
  if err != nil {
    log.Fatal(err)
  }

  log.Println("Verify Now: ", vc)

  //String representation
  log.Println(t)
}
