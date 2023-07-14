package crypt

// 计算merkle的root
func MerkleRoot(content [][]byte) [32]byte {
	if len(content)%2 == 1 {
		content = append(content, content[len(content)-1])
	}

	length := len(content)
	if length == 1 {
		return DoubleSha256(content[0])
	}

	if length == 2 {
		root := append(content[0], content[1]...)
		return DoubleSha256(root)
	}

	weight := 1
	for weight < length/2 {
		weight = weight * 2
	}

	left := content[0:weight]
	right := content[weight:]
	root := []byte{}
	for _, b := range MerkleRoot(left) {
		root = append(root, b)
	}
	for _, b := range MerkleRoot(right) {
		root = append(root, b)
	}
	return DoubleSha256(root)
}

// /*
// 计算merkle的root
// */
// func MerkleRoot(content ...[]byte) []byte {
// 	//先克隆
// 	contentCopy := make([][]byte, len(content))
// 	copy(contentCopy, content)

// 	return merkleRoot(contentCopy)
// }
