package main

import (
	"math/rand"
	"testing"
	// Import your Red-Black Tree package
)

func TestRedBlackTreeInsertion(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Normal Case 1",
			input: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "Normal Case 2",
			input: []int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:  "Edge Case 1 (Empty Slice)",
			input: []int{},
		},
		{
			name:  "Edge Case 2 (Single Element)",
			input: []int{42},
		},
		{
			name:  "Edge Case 3 (Duplicates)",
			input: []int{2, 3, 4, 3, 1, 2, 5, 5, 4, 1},
		},
		{
			name:  "Edge Case 4 (Negative Integers)",
			input: []int{-5, -2, -10, -1, -7, -3, -8},
		},
		{
			name:  "Edge Case 5 (All Identical Elements)",
			input: []int{8, 8, 8, 8, 8, 8, 8},
		},
		{
			name:  "Edge Case 6 (Random Integers Big Range)",
			input: []int{560, 746, 74, 9, 754, 231, 880, 797, 131, 478, 435, 111, 541, 816, 833, 750, 703, 522, 965, 310, 106, 869, 677, 962, 292, 325, 67, 184, 733, 342, 832, 19, 287, 335, 168, 809, 983, 506, 88, 761, 395, 492, 88, 610, 775, 782, 992, 414, 641, 911, 67, 92, 810, 926, 22, 316, 329, 577, 331, 554, 337, 356, 915, 540, 125, 623, 844, 450, 649, 212, 259, 6, 329, 790, 494, 775, 763, 836, 771, 183, 44, 833, 108, 700, 853, 170, 326, 631, 673, 938, 943, 242, 163, 128, 849, 325, 19, 876, 248, 187, 420, 364, 746, 472, 534, 33, 378, 429, 888, 297, 431, 249, 292, 446, 162, 347, 253, 98, 402, 704, 678, 499, 802, 633, 284, 14, 769, 393, 231, 610, 610, 178, 736, 886, 448, 680, 982, 972, 250, 354, 918, 266, 171, 688, 484, 630, 351, 642, 172, 73, 834, 499, 837, 466, 937, 926, 340, 152, 365, 673, 365, 304, 80, 66, 92, 870, 559, 641, 338, 797, 294, 284, 366, 595, 732, 256, 606, 674, 508, 366, 715, 834, 209, 183, 969, 555, 20, 841, 459, 169, 276, 722, 726, 326, 228, 216, 575, 141, 836, 84, 340, 873, 197, 971, 623, 884, 147, 984, 899, 487, 641, 34, 501, 551, 286, 814, 935, 301, 822, 788, 918, 532, 901, 387, 837, 41, 378, 1, 153, 85, 176, 305, 955, 601, 94, 370, 611, 165, 868, 8, 603, 975, 628, 270, 564, 960, 128, 414, 236, 611, 151, 547, 177, 953, 403, 536, 462, 308, 236, 881, 798, 779, 537, 475, 253, 200, 300, 274, 797, 107, 766, 616, 231, 743, 139, 347, 888, 29, 336, 354, 73, 320, 722, 795, 648, 283, 147, 162, 734, 565, 978, 708, 177, 620, 823, 441, 475, 196, 836, 381, 363, 536, 17, 648, 212, 991, 971, 238, 132, 7, 12, 860, 208, 491, 424, 140, 657, 695, 631, 895, 975, 677, 827, 115, 155, 817, 698, 405, 749, 412, 880, 395, 534, 629, 665, 146, 965, 268, 122, 848, 705, 14, 292, 918, 0, 653, 600, 474, 525, 66, 779, 708, 650, 430, 854, 495, 575, 571, 959, 622, 904, 519, 78, 486, 295, 676, 780, 135, 915, 731, 564, 955, 955, 533, 958, 376, 862, 124, 856, 146, 585, 487, 996, 145, 55, 582, 269, 944, 250, 121, 671, 123, 820, 974, 362, 892, 450, 963, 92, 168, 665, 352, 437, 794, 569, 196, 577, 351, 418, 805, 474, 482, 86, 838, 857, 88, 327, 240, 196, 900, 688, 44, 847},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test for %s", tc.name)
			// Create an empty Red-Black Tree
			tree := GenerateTree() // Use your actual constructor function

			// Insert elements from the current test case's input slice
			for _, value := range tc.input {
				tree.Insert(value) // Replace with your insertion function
			}

			// Verify the Red-Black Tree properties
			if !verify_rbt_properties(tree) {
				t.Errorf("Red-Black Tree properties violated after insertions")
			} else {
				t.Logf("Red-Black Tree properties passed for %s", tc.name)
			}

			// Add more assertions to test the structure or specific properties of the tree.
		})
	}
}

func TestRedBlackTreeInsertionRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		// Generate a random slice of integers with random size
		size := rand.Intn(1000)
		input := make([]int, size)
		for j := 0; j < size; j++ {
			input[j] = rand.Intn(1000)
		}

		// fmt.Println("Testing insertion of", input)

		// Insert the integers into the Red-Black Tree
		tree := GenerateTree()
		for _, x := range input {
			tree.Insert(x)
		}

		// Check that the tree is valid
		if !verify_rbt_properties(tree) {
			t.Errorf("Tree is not valid after insertion of %v", input)
		}
	}
}

func benchmarkInsert(n int, b *testing.B) {
	// Create an empty Red-Black Tree
	tree := GenerateTree()

	// Generate n random integers
	input := rand.Perm(n)

	// Reset the timer
	b.ResetTimer()

	// Insert the integers into the tree
	for i := 0; i < b.N; i++ {
		for _, x := range input {
			tree.Insert(x)
		}
	}
}

func BenchmarkInsert100(b *testing.B)   { benchmarkInsert(100, b) }
func BenchmarkInsert1000(b *testing.B)  { benchmarkInsert(1000, b) }
func BenchmarkInsert10000(b *testing.B) { benchmarkInsert(10000, b) }
