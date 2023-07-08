package crypt

//计算merkle的root
func MerkleRoot(content [][]byte) []byte {
	if len(content)%2 == 1 {
		content = append(content, content[len(content)-1])
	}

	length := len(content)
	if length == 1 {
		return Sha256(Sha256(content[0]))
	}

	if length == 2 {
		root := append(content[0], content[1]...)
		return Sha256(Sha256(root))
	}

	weight := 1
	for weight < length/2 {
		weight = weight * 2
	}

	left := content[0:weight]
	right := content[weight:]
	root := append(MerkleRoot(left), MerkleRoot(right)...)
	return Sha256(Sha256(root))
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
