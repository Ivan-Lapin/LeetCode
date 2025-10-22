# LeetCode Solutions in Go

This repository contains a collection of solutions to LeetCode problems implemented in the Go programming language. The solutions cover a variety of algorithmic and data structure challenges, ranging from easy to medium difficulty, and are designed to help improve problem-solving skills and proficiency in Go.

## Project Overview

The code is organized in a single `main.go` file, with each function corresponding to a specific LeetCode problem. The solutions leverage Go's standard library for tasks such as string manipulation, sorting, and mathematical operations. This project serves as a reference for learning algorithmic problem-solving in Go and can be extended by adding more solutions or optimizing existing ones.

## LeetCode Problems

The following table lists the implemented LeetCode problems, including their problem number, title, a brief description of the task, and difficulty level as per LeetCode's classification.

| #   |                         Title                     |                      Description                           | Difficulty |
|-----|---------------------------------------------------|------------------------------------------------------------|------------|
| 1   | [Two Sum](https://leetcode.com/problems/two-sum/) | Given an array of integers `nums` and a target integer `target`, find the indices of two numbers that add up to `target`. | Easy |
| 9   | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | Determine if an integer `x` is a palindrome without converting it to a string. | Easy |
| 13  | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | Convert a Roman numeral string to an integer. | Easy |
| 14  | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | Find the longest common prefix among an array of strings. | Easy |
| 20  | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | Check if a string containing parentheses is valid (properly nested). | Easy |
| 21  | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | Merge two sorted linked lists into a new sorted linked list. | Easy |
| 26  | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | Remove duplicates from a sorted array in-place and return the new length. | Easy |
| 27  | [Remove Element](https://leetcode.com/problems/remove-element/) | Remove all occurrences of a value from an array in-place and return the new length. | Easy |
| 28  | [Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | Find the index of the first occurrence of `needle` in `haystack`, or return -1 if not found. | Easy |
| 35  | [Search Insert Position](https://leetcode.com/problems/search-insert-position/) | Find the index where a target value should be inserted in a sorted array to maintain order. | Easy |
| 58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/) | Return the length of the last word in a string. | Easy |
| 66  | [Plus One](https://leetcode.com/problems/plus-one/) | Increment a number represented as an array of digits by one. | Easy |
| 67  | [Add Binary](https://leetcode.com/problems/add-binary/) | Add two binary strings and return the result as a binary string. | Easy |
| 69  | [Sqrt(x)](https://leetcode.com/problems/sqrtx/) | Compute the integer square root of a non-negative integer `x`. | Easy |
| 70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | Find the number of distinct ways to climb `n` stairs, taking 1 or 2 steps at a time. | Easy |
| 83  | [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | Remove duplicates from a sorted linked list. | Easy |
| 118 | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/) | Generate the first `numRows` of Pascal's triangle. | Easy |
| 119 | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/) | Return the `rowIndex`-th row of Pascal's triangle. | Easy |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Find the maximum profit from buying and selling a stock given an array of prices. | Easy |
| 125 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | Check if a string is a palindrome, ignoring non-alphanumeric characters and case. | Easy |
| 136 | [Single Number](https://leetcode.com/problems/single-number/) | Find the only number in an array that appears exactly once, while others appear twice. | Easy |
| 169 | [Majority Element](https://leetcode.com/problems/majority-element/) | Find the element that appears more than `n/2` times in an array. | Easy |
| 190 | [Reverse Bits](https://leetcode.com/problems/reverse-bits/) | Reverse the bits of a 32-bit unsigned integer. | Easy |
| 191 | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) | Count the number of 1 bits in an unsigned integer (Hamming weight). | Easy |
| 202 | [Happy Number](https://leetcode.com/problems/happy-number/) | Determine if a number is "happy" by checking if the sum of squares of its digits eventually reaches 1. | Easy |
| 205 | [Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/) | Check if two strings are isomorphic (characters map uniquely to each other). | Easy |
| 219 | [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/) | Check if there are duplicates within `k` indices in an array. | Easy |
| 228 | [Summary Ranges](https://leetcode.com/problems/summary-ranges/) | Summarize consecutive ranges in a sorted array as strings (e.g., "0->2"). | Easy |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Check if two strings are anagrams of each other. | Easy |
| 290 | [Word Pattern](https://leetcode.com/problems/word-pattern/) | Check if a pattern matches a string of words with a one-to-one mapping. | Easy |
| 303 | [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/) | Compute the sum of elements in a given range of an array using prefix sums. | Easy |
| 344 | [Reverse String](https://leetcode.com/problems/reverse-string/) | Reverse a string in-place. | Easy |
| 349 | [Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/) | Find the intersection of two arrays, returning unique elements. | Easy |
| 387 | [First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/) | Find the index of the first non-repeating character in a string. | Easy |
| 389 | [Find the Difference](https://leetcode.com/problems/find-the-difference/) | Find the extra character in string `t` compared to string `s`. | Easy |
| 392 | [Is Subsequence](https://leetcode.com/problems/is-subsequence/) | Check if string `s` is a subsequence of string `t`. | Easy |
| 409 | [Longest Palindrome](https://leetcode.com/problems/longest-palindrome/) | Find the length of the longest palindrome that can be formed from a string. | Easy |
| 415 | [Add Strings](https://leetcode.com/problems/add-strings/) | Add two numbers represented as strings without using built-in conversion to integers. | Easy |
| 459 | [Repeated Substring Pattern](https://leetcode.com/problems/repeated-substring-pattern/) | Check if a string can be formed by repeating a substring. | Easy |
| 476 | [Number Complement](https://leetcode.com/problems/number-complement/) | Find the binary complement of a given integer. | Easy |
| 496 | [Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/) | Find the next greater element for each element in `nums1` within `nums2`. | Easy |
| 507 | [Perfect Number](https://leetcode.com/problems/perfect-number/) | Check if a number is perfect (sum of proper divisors equals the number). | Easy |
| 520 | [Detect Capital](https://leetcode.com/problems/detect-capital/) | Check if a word follows valid capitalization rules. | Easy |
| 557 | [Reverse Words in a String III](https://leetcode.com/problems/reverse-words-in-a-string-iii/) | Reverse each word in a string while maintaining word order. | Easy |
| 561 | [Array Partition I](https://leetcode.com/problems/array-partition-i/) | Pair elements in an array to maximize the sum of the minimums of each pair. | Easy |
| 575 | [Distribute Candies](https://leetcode.com/problems/distribute-candies/) | Determine the maximum number of candy types a sister can get when candies are divided equally. | Easy |
| 599 | [Minimum Index Sum of Two Lists](https://leetcode.com/problems/minimum-index-sum-of-two-lists/) | Find common restaurant names with the minimum sum of indices from two lists. | Easy |
| 605 | [Can Place Flowers](https://leetcode.com/problems/can-place-flowers/) | Check if `n` flowers can be planted in a flowerbed without adjacent plants. | Easy |
| 606 | [Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/) | Convert a binary tree to a string representation using preorder traversal. | Easy |
| 643 | [Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/) | Find the maximum average of a contiguous subarray of length `k`. | Easy |
| 645 | [Set Mismatch](https://leetcode.com/problems/set-mismatch/) | Find the duplicated and missing numbers in an array. | Easy |
| 657 | [Robot Return to Origin](https://leetcode.com/problems/robot-return-to-origin/) | Check if a robot returns to the origin after a sequence of moves. | Easy |
| 674 | [Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/) | Find the length of the longest continuous increasing subsequence. | Easy |
| 680 | [Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/) | Check if a string can be a palindrome by removing at most one character. | Easy |
| 682 | [Baseball Game](https://leetcode.com/problems/baseball-game/) | Compute the total score of a baseball game given a list of operations. | Easy |
| 693 | [Binary Number with Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/) | Check if a number's binary representation has alternating bits. | Easy |
| 696 | [Count Binary Substrings](https://leetcode.com/problems/count-binary-substrings/) | Count the number of substrings with equal consecutive 0s and 1s. | Easy |
| 697 | [Degree of an Array](https://leetcode.com/problems/degree-of-an-array/) | Find the shortest subarray containing all occurrences of the most frequent element. | Easy |
| 700 | [Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/) | Search for a value in a binary search tree. | Easy |
| 704 | [Binary Search](https://leetcode.com/problems/binary-search/) | Find a target value in a sorted array using binary search. | Easy |
| 709 | [To Lower Case](https://leetcode.com/problems/to-lower-case/) | Convert a string to lowercase. | Easy |
| 717 | [1-bit and 2-bit Characters](https://leetcode.com/problems/1-bit-and-2-bit-characters/) | Check if the last character in a binary array is a 1-bit character. | Easy |
| 724 | [Find Pivot Index](https://leetcode.com/problems/find-pivot-index/) | Find the index where the sum of elements on the left equals the sum on the right. | Easy |
| 728 | [Self Dividing Numbers](https://leetcode.com/problems/self-dividing-numbers/) | Find all self-dividing numbers in a given range. | Easy |
| 744 | [Find Smallest Letter Greater Than Target](https://leetcode.com/problems/find-smallest-letter-greater-than-target/) | Find the smallest letter greater than a target in a sorted array of letters. | Easy |
| 746 | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) | Find the minimum cost to reach the top of a staircase with given costs per step. | Easy |
| 747 | [Largest Number At Least Twice of Others](https://leetcode.com/problems/largest-number-at-least-twice-of-others/) | Check if the largest number is at least twice as large as every other number. | Easy |
| 748 | [Shortest Completing Word](https://leetcode.com/problems/shortest-completing-word/) | Find the shortest word that contains all letters from a license plate. | Easy |
| 766 | [Toeplitz Matrix](https://leetcode.com/problems/toeplitz-matrix/) | Check if a matrix is a Toeplitz matrix (diagonals have the same elements). | Easy |
| 771 | [Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/) | Count how many stones are jewels given two strings. | Easy |
| 783 | [Minimum Distance Between BST Nodes](https://leetcode.com/problems/minimum-distance-between-bst-nodes/) | Find the minimum difference between any two nodes in a BST. | Easy |
| 796 | [Rotate String](https://leetcode.com/problems/rotate-string/) | Check if one string is a rotation of another. | Easy |
| 804 | [Unique Morse Code Words](https://leetcode.com/problems/unique-morse-code-words/) | Count unique Morse code representations of a list of words. | Easy |
| 806 | [Number of Lines To Write String](https://leetcode.com/problems/number-of-lines-to-write-string/) | Calculate the number of lines and pixels needed to write a string. | Easy |
| 821 | [Shortest Distance to a Character](https://leetcode.com/problems/shortest-distance-to-a-character/) | Find the shortest distance from each character in a string to a target character. | Easy |
| 824 | [Goat Latin](https://leetcode.com/problems/goat-latin/) | Convert a sentence to Goat Latin by applying specific rules to each word. | Easy |
| 830 | [Positions of Large Groups](https://leetcode.com/problems/positions-of-large-groups/) | Find positions of groups of 3 or more identical characters in a string. | Easy |
| 832 | [Flipping an Image](https://leetcode.com/problems/flipping-an-image/) | Flip a binary matrix horizontally and invert it. | Easy |
| 836 | [Rectangle Overlap](https://leetcode.com/problems/rectangle-overlap/) | Check if two rectangles overlap. | Easy |
| 844 | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/) | Compare two strings after processing backspace characters. | Easy |
| 852 | [Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/) | Find the peak index in a mountain array (increases then decreases). | Easy |
| 859 | [Buddy Strings](https://leetcode.com.com/problems/buddy-strings/) | Check if two strings can be made equal by swapping two characters. | Easy |
| 860 | [Lemonade Change](https://leetcode.com/problems/lemonade-change/) | Determine if correct change can be provided for lemonade purchases. | Easy |
| 867 | [Transpose Matrix](https://leetcode.com/problems/transpose-matrix/) | Transpose a matrix (swap rows and columns). | Easy |
| 872 | [Leaf-Similar Trees](https://leetcode.com/problems/leaf-similar-trees/) | Check if two binary trees have the same leaf sequence. | Easy |
| 876 | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) | Find the middle node of a linked list. | Easy |
| 883 | [Projection Area of 3D Shapes](https://leetcode.com/problems/projection-area-of-3d-shapes/) | Calculate the projection area of a 3D shape on three planes. | Easy |
| 884 | [Uncommon Words from Two Sentences](https://leetcode.com/problems/uncommon-words-from-two-sentences/) | Find words that appear exactly once in two sentences. | Easy |
| 888 | [Fair Candy Swap](https://leetcode.com/problems/fair-candy-swap/) | Find a pair of candies to swap so both children have the same total. | Easy |
| 896 | [Monotonic Array](https://leetcode.com/problems/monotonic-array/) | Check if an array is monotonic (either increasing or decreasing). | Easy |
| 897 | [Increasing Order Search Tree](https://leetcode.com/problems/increasing-order-search-tree/) | Convert a BST to a sorted linked list in increasing order. | Easy |
| 905 | [Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/) | Sort an array so all even numbers come before odd numbers. | Easy |
| 908 | [Smallest Range I](https://leetcode.com/problems/smallest-range-i/) | Find the smallest possible range after adding a value within [-K, K] to each element. | Easy |
| 914 | [X of a Kind in a Deck of Cards](https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/) | Check if a deck can be divided into groups of the same size. | Easy |
| 917 | [Reverse Only Letters](https://leetcode.com/problems/reverse-only-letters/) | Reverse only the letters in a string, keeping non-letters in place. | Easy |
| 922 | [Sort Array By Parity II](https://leetcode.com/problems/sort-array-by-parity-ii/) | Sort an array so even-indexed elements are even and odd-indexed are odd. | Easy |
| 925 | [Long Pressed Name](https://leetcode.com/problems/long-pressed-name/) | Check if a typed name is a valid long-pressed version of the original name. | Easy |
| 929 | [Unique Email Addresses](https://leetcode.com/problems/unique-email-addresses/) | Count unique email addresses after applying normalization rules. | Easy |
| 933 | [Number of Recent Calls](https://leetcode.com/problems/number-of-recent-calls/) | Track the number of pings in the last 3000 milliseconds. | Easy |
| 938 | [Range Sum of BST](https://leetcode.com/problems/range-sum-of-bst/) | Find the sum of all nodes in a BST within a given range. | Easy |
| 941 | [Valid Mountain Array](https://leetcode.com/problems/valid-mountain-array/) | Check if an array forms a valid mountain (strictly increasing then decreasing). | Easy |
| 942 | [DI String Match](https://leetcode.com/problems/di-string-match/) | Construct an array that matches a string of 'I' (increase) and 'D' (decrease). | Easy |
| 944 | [Delete Columns to Make Sorted](https://leetcode.com/problems/delete-columns-to-make-sorted/) | Find the number of columns to delete to make a matrix sorted. | Easy |
| 953 | [Verifying an Alien Dictionary](https://leetcode.com/problems/verifying-an-alien-dictionary/) | Check if words are sorted in a custom alien dictionary order. | Easy |
| 961 | [N-Repeated Element in Size 2N Array](https://leetcode.com/problems/n-repeated-element-in-size-2n-array/) | Find the element that appears `N` times in an array of size `2N`. | Easy |
| 965 | [Univalued Binary Tree](https://leetcode.com/problems/univalued-binary-tree/) | Check if all nodes in a binary tree have the same value. | Easy |
| 977 | [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/) | Return the squares of a sorted array in sorted order. | Easy |
| 989 | [Add to Array-Form of Integer](https://leetcode.com/problems/add-to-array-form-of-integer/) | Add an integer to a number represented as an array of digits. | Easy |
| 993 | [Cousins in Binary Tree](https://leetcode.com/problems/cousins-in-binary-tree/) | Check if two nodes in a binary tree are cousins (same depth, different parents). | Easy |
| 997 | [Find the Town Judge](https://leetcode.com/problems/find-the-town-judge/) | Find the town judge who is trusted by everyone but trusts no one. | Easy |
| 999 | [Available Captures for Rook](https://leetcode.com/problems/available-captures-for-rook/) | Count the number of pawns a rook can capture in one move. | Easy |
| 1002 | [Find Common Characters](https://leetcode.com/problems/find-common-characters/) | Find all characters that appear in every string in a list. | Easy |
| 1005 | [Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/) | Maximize the sum of an array by negating up to `K` elements. | Easy |
| 1009 | [Complement of Base 10 Integer](https://leetcode.com/problems/complement-of-base-10-integer/) | Find the bitwise complement of a base-10 integer. | Easy |
| 1013 | [Partition Array Into Three Parts With Equal Sum](https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/) | Check if an array can be partitioned into three parts with equal sums. | Easy |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/) | Check which prefixes of a binary array are divisible by 5. | Easy |
| 1021 | [Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/) | Remove the outermost parentheses from a valid parentheses string. | Easy |
| 1025 | [Divisor Game](https://leetcode.com/problems/divisor-game/) | Determine if Alice wins a divisor game against Bob. | Easy |
| 1037 | [Valid Boomerang](https://leetcode.com/problems/valid-boomerang/) | Check if three points form a valid boomerang (not collinear). | Easy |
| 1046 | [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/) | Simulate smashing stones until at most one remains and return its weight. | Easy |
| 1047 | [Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) | Remove all adjacent duplicate characters from a string. | Easy |
| 1051 | [Height Checker](https://leetcode.com/problems/height-checker/) | Count mismatches between a height array and its sorted version. | Easy |
| 1078 | [Occurrences After Bigram](https://leetcode.com/problems/occurrences-after-bigram/) | Find all words that appear after a specific bigram in a text. | Easy |
| 1089 | [Duplicate Zeros](https://leetcode.com/problems/duplicate-zeros/) | Duplicate each zero in an array, shifting other elements to the right. | Easy |
| 1108 | [Defanging an IP Address](https://leetcode.com/problems/defanging-an-ip-address/) | Replace periods in an IP address with `[.]`. | Easy |
| 1122 | [Relative Sort Array](https://leetcode.com/problems/relative-sort-array/) | Sort `arr1` based on the order in `arr2`, with remaining elements sorted. | Easy |
| 1137 | [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/) | Find the `n`-th Tribonacci number, where each number is the sum of the three preceding ones. | Easy |
| 1154 | [Day of the Year](https://leetcode.com/problems/day-of-the-year/) | Calculate the day of the year for a given date. | Easy |
| 1160 | [Find Words That Can Be Formed by Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/) | Find words that can be formed using a given set of characters. | Easy |
| 1185 | [Day of the Week](https://leetcode.com/problems/day-of-the-week/) | Determine the day of the week for a given date. | Easy |
| 1189 | [Maximum Number of Balloons](https://leetcode.com/problems/maximum-number-of-balloons/) | Find the maximum number of times "balloon" can be formed from a string. | Easy |
| 1207 | [Unique Number of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) | Check if the frequency of each number in an array is unique. | Easy |
| 1221 | [Split a String in Balanced Strings](https://leetcode.com/problems/split-a-string-in-balanced-strings/) | Split a string into the maximum number of balanced substrings. | Easy |
| 1232 | [Check If It Is a Straight Line](https://leetcode.com/problems/check-if-it-is-a-straight-line/) | Check if all points lie on a straight line. | Easy |
| 1252 | [Cells with Odd Values in a Matrix](https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/) | Count cells with odd values after applying operations to a matrix. | Easy |
| 1266 | [Minimum Time Visiting All Points](https://leetcode.com/problems/minimum-time-visiting-all-points/) | Calculate the minimum time to visit all points in a 2D plane. | Easy |
| 1281 | [Subtract the Product and Sum of Digits of an Integer](https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/) | Compute the difference between the product and sum of an integer's digits. | Easy |
| 1290 | [Convert Binary Number in a Linked List to Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/) | Convert a binary number represented as a linked list to an integer. | Easy |
| 1295 | [Find Numbers with Even Number of Digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/) | Count numbers with an even number of digits in an array. | Easy |
| 1299 | [Replace Elements with Greatest Element on Right Side](https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/) | Replace each element with the greatest element to its right. | Easy |
| 1304 | [Find N Unique Integers Sum up to Zero](https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/) | Generate an array of `n` unique integers that sum to zero. | Easy |
| 1309 | [Decrypt String from Alphabet to Integer Mapping](https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/) | Decrypt a string where numbers map to letters (1->a, 2->b, ..., 26->z). | Easy |
| 1313 | [Decompress Run-Length Encoded List](https://leetcode.com/problems/decompress-run-length-encoded-list/) | Decompress a run-length encoded list into a regular list. | Easy |
| 1323 | [Maximum 69 Number](https://leetcode.com/problems/maximum-69-number/) | Maximize a number by changing at most one digit (6 to 9 or 9 to 6). | Easy |
| 1332 | [Remove Palindromic Subsequences](https://leetcode.com/problems/remove-palindromic-subsequences/) | Find the minimum number of steps to remove all characters in a string using palindromic subsequences. | Easy |
| 1342 | [Number of Steps to Reduce a Number to Zero](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/) | Count steps to reduce a number to zero by dividing by 2 or subtracting 1. | Easy |
| 1351 | [Count Negative Numbers in a Sorted Matrix](https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/) | Count negative numbers in a sorted matrix. | Easy |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/) | Sort an array based on the number of 1 bits in each integer. | Easy |
| 1365 | [How Many Numbers Are Smaller Than the Current Number](https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/) | Count how many numbers are smaller than each number in an array. | Easy |
| 1374 | [Generate a String With Characters That Have Odd Counts](https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/) | Generate a string of length `n` where each character has an odd count. | Easy |
| 1389 | [Create Target Array in the Given Order](https://leetcode.com/problems/create-target-array-in-the-given-order/) | Create a target array by inserting elements at given indices. | Easy |
| 1403 | [Minimum Subsequence in Non-Increasing Order](https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/) | Find the minimum subsequence with a sum greater than the rest of the array. | Easy |
| 1413 | [Minimum Value to Get Positive Step by Step Sum](https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/) | Find the minimum starting value to ensure the step-by-step sum is always positive. | Easy |
| 1417 | [Reformat The String](https://leetcode.com/problems/reformat-the-string/) | Reformat a string by interleaving digits and letters. | Easy |
| 1422 | [Maximum Score After Splitting a String](https://leetcode.com/problems/maximum-score-after-splitting-a-string/) | Find the maximum score by splitting a string into two parts. | Easy |
| 1431 | [Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) | Check if kids can have the most candies after receiving extra candies. | Easy |
| 1436 | [Destination City](https://leetcode.com/problems/destination-city/) | Find the destination city in a list of paths (city with no outgoing path). | Easy |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/) | Build a target array using push and pop operations. | Easy |
| 1450 | [Number of Students Doing Homework at a Given Time](https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/) | Count students doing homework at a specific time. | Easy |
| 1455 | [Check If a Word Occurs As a Prefix of Any Word in a Sentence](https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/) | Check if a word is a prefix of any word in a sentence. | Easy |
| 1460 | [Make Two Arrays Equal by Reversing Sub-arrays](https://leetcode.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/) | Check if two arrays can be made equal by reversing subarrays. | Easy |
| 1464 | [Maximum Product of Two Elements in an Array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/) | Find the maximum product of two numbers minus one in an array. | Easy |
| 1470 | [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/) | Shuffle an array by interleaving the first and second halves. | Easy |
| 1475 | [Final Prices With a Special Discount in a Shop](https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/) | Compute final prices after applying discounts based on later items. | Easy |
| 1480 | [Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/) | Compute the running sum of an array. | Easy |
| 1486 | [XOR Operation in an Array](https://leetcode.com/problems/xor-operation-in-an-array/) | Compute the XOR of all elements in an array defined by a formula. | Easy |
| 1491 | [Average Salary Excluding the Minimum and Maximum Salary](https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/) | Compute the average salary excluding the minimum and maximum. | Easy |
| 1496 | [Path Crossing](https://leetcode.com/problems/path-crossing/) | Check if a path crosses itself on a 2D plane. | Easy |
| 1502 | [Can Make Arithmetic Progression From Sequence](https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/) | Check if an array can be rearranged into an arithmetic progression. | Easy |
| 1507 | [Reformat Date](https://leetcode.com/problems/reformat-date/) | Reformat a date string from a specific format to `YYYY-MM-DD`. | Easy |
| 1512 | [Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/) | Count pairs of equal elements in an array. | Easy |
| 1518 | [Water Bottles](https://leetcode.com/problems/water-bottles/) | Calculate the total number of water bottles that can be drunk. | Easy |
| 1523 | [Count Odd Numbers in an Interval Range](https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/) | Count odd numbers in a given range. | Easy |
| 1528 | [Shuffle String](https://leetcode.com/problems/shuffle-string/) | Restore a shuffled string using given indices. | Easy |
| 1534 | [Count Good Triplets](https://leetcode.com/problems/count-good-triplets/) | Count triplets in an array that satisfy given conditions. | Easy |
| 1539 | [Kth Missing Positive Number](https://leetcode.com/problems/kth-missing-positive-number/) | Find the `k`-th missing positive integer in a sorted array. | Easy |
| 1544 | [Make The String Great](https://leetcode.com/problems/make-the-string-great/) | Remove adjacent characters that differ only by case to make a string "great". | Easy |
| 1550 | [Three Consecutive Odds](https://leetcode.com/problems/three-consecutive-odds/) | Check if there are three consecutive odd numbers in an array. | Easy |
| 1556 | [Thousand Separator](https://leetcode.com/problems/thousand-separator/) | Add thousand separators to a number as a string. | Easy |
| 1560 | [Most Visited Sector in a Circular Track](https://leetcode.com/problems/most-visited-sector-in-a-circular-track/) | Find the most visited sectors in a circular track. | Easy |
| 1572 | [Matrix Diagonal Sum](https://leetcode.com/problems/matrix-diagonal-sum/) | Compute the sum of elements on the primary and secondary diagonals of a matrix. | Easy |
| 1582 | [Special Positions in a Binary Matrix](https://leetcode.com/problems/special-positions-in-a-binary-matrix/) | Count special positions in a binary matrix (only 1 in row and column). | Easy |
| 1588 | [Sum of All Odd Length Subarrays](https://leetcode.com/problems/sum-of-all-odd-length-subarrays/) | Compute the sum of all odd-length subarrays in an array. | Easy |
| 1592 | [Rearrange Spaces Between Words](https://leetcode.com/problems/rearrange-spaces-between-words/) | Rearrange spaces in a string to have equal spaces between words. | Easy |
| 1598 | [Crawler Log Folder](https://leetcode.com/problems/crawler-log-folder/) | Compute the minimum operations to return to the main folder after a series of logs. | Easy |
| 1603 | [Design Parking System](https://leetcode.com/problems/design-parking-system/) | Design a parking system to track available spaces for different car types. | Easy |
| 1608 | [Special Array With X Elements Greater Than or Equal X](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/) | Find if an array is special (has `x` elements greater than or equal to `x`). | Easy |
| 1614 | [Maximum Nesting Depth of the Parentheses](https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/) | Find the maximum nesting depth of parentheses in a string. | Easy |
| 1619 | [Mean of Array After Removing Some Elements](https://leetcode.com/problems/mean-of-array-after-removing-some-elements/) | Compute the mean of an array after removing the smallest and largest 5% of elements. | Easy |
| 1624 | [Largest Substring Between Two Equal Characters](https://leetcode.com/problems/largest-substring-between-two-equal-characters/) | Find the length of the largest substring between two equal characters. | Easy |
| 1629 | [Slowest Key](https://leetcode.com/problems/slowest-key/) | Find the key with the longest duration in a sequence of key presses. | Easy |
| 1636 | [Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/) | Sort an array by increasing frequency and then by decreasing value. | Easy |
| 1640 | [Check Array Formation Through Concatenation](https://leetcode.com/problems/check-array-formation-through-concatenation/) | Check if an array can be formed by concatenating given pieces. | Easy |
| 1646 | [Get Maximum in Generated Array](https://leetcode.com/problems/get-maximum-in-generated-array/) | Find the maximum value in an array generated by specific rules. | Easy |
| 1652 | [Defuse the Bomb](https://leetcode.com/problems/defuse-the-bomb/) | Decrypt a circular array by summing elements in a given range. | Easy |
| 1656 | [Design an Ordered Stream](https://leetcode.com/problems/design-an-ordered-stream/) | Design a stream that returns values in order as they are inserted. | Easy |
| 1662 | [Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/) | Check if two string arrays represent the same string when concatenated. | Easy |
| 1672 | [Richest Customer Wealth](https://leetcode.com/problems/richest-customer-wealth/) | Find the maximum wealth of any customer given their bank accounts. | Easy |
| 1678 | [Goal Parser Interpretation](https://leetcode.com/problems/goal-parser-interpretation/) | Interpret a string command to produce the corresponding output. | Easy |
| 1684 | [Count the Number of Consistent Strings](https://leetcode.com/problems/count-the-number-of-consistent-strings/) | Count strings that only contain characters from a given set. | Easy |
| 1688 | [Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/) | Count the number of matches in a tournament with `n` teams. | Easy |
| 1694 | [Reformat Phone Number](https://leetcode.com/problems/reformat-phone-number/) | Reformat a phone number to group digits in blocks of 2 or 3. | Easy |
| 1704 | [Determine if String Halves Are Alike](https://leetcode.com/problems/determine-if-string-halves-are-alike/) | Check if two halves of a string have the same number of vowels. | Easy |
| 1710 | [Maximum Units on a Truck](https://leetcode.com/problems/maximum-units-on-a-truck/) | Maximize the total units on a truck given box types and sizes. | Easy |
| 1716 | [Calculate Money in Leetcode Bank](https://leetcode.com/problems/calculate-money-in-leetcode-bank/) | Calculate the total money in a bank after `n` days with a weekly pattern. | Easy |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/) | Decode an array where each element is the XOR of the previous and current elements. | Easy |
| 1725 | [Number Of Rectangles That Can Form The Largest Square](https://leetcode.com/problems/number-of-rectangles-that-can-form-the-largest-square/) | Count rectangles that can form the largest possible square. | Easy |
| 1732 | [Find the Highest Altitude](https://leetcode.com/problems/find-the-highest-altitude/) | Find the highest altitude given an array of altitude gains. | Easy |
| 1742 | [Maximum Number of Balls in a Box](https://leetcode.com/problems/maximum-number-of-balls-in-a-box/) | Find the box with the maximum number of balls based on digit sums. | Easy |
| 1748 | [Sum of Unique Elements](https://leetcode.com/problems/sum-of-unique-elements/) | Compute the sum of elements that appear exactly once in an array. | Easy |
| 1752 | [Check if Array Is Sorted and Rotated](https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/) | Check if an array is sorted and possibly rotated. | Easy |
| 1758 | [Minimum Changes To Make Alternating Binary String](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/) | Find the minimum changes to make a binary string alternating. | Easy |
| 1768 | [Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) | Merge two strings by alternating characters. | Easy |
| 1773 | [Count Items Matching a Rule](https://leetcode.com/problems/count-items-matching-a-rule/) | Count items that match a given rule based on type, color, or name. | Easy |
| 1784 | [Check if Binary String Has at Most One Segment of Ones](https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/) | Check if a binary string has at most one continuous segment of ones. | Easy |
| 1790 | [Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/) | Check if two strings can be made equal by swapping at most one pair of characters. | Easy |
| 1796 | [Second Largest Digit in a String](https://leetcode.com/problems/second-largest-digit-in-a-string/) | Find the second largest digit in a string. | Easy |
| 1800 | [Maximum Ascending Subarray Sum](https://leetcode.com/problems/maximum-ascending-subarray-sum/) | Find the maximum sum of a contiguous ascending subarray. | Easy |
| 1805 | [Number of Different Integers in a String](https://leetcode.com/problems/number-of-different-integers-in-a-string/) | Count distinct integers in a string. | Easy |
| 1812 | [Determine Color of a Chessboard Square](https://leetcode.com/problems/determine-color-of-a-chessboard-square/) | Determine if a chessboard square is white or black. | Easy |
| 1816 | [Truncate Sentence](https://leetcode.com/problems/truncate-sentence/) | Truncate a sentence to contain only the first `k` words. | Easy |
| 1822 | [Sign of the Product of an Array](https://leetcode.com/problems/sign-of-the-product-of-an-array/) | Determine the sign of the product of all elements in an array. | Easy |
| 1827 | [Minimum Operations to Make the Array Increasing](https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/) | Calculate the minimum operations to make an array strictly increasing. | Easy |
| 1832 | [Check if the Sentence Is Pangram](https://leetcode.com/problems/check-if-the-sentence-is-pangram/) | Check if a sentence contains all letters of the alphabet. | Easy |
| 1837 | [Sum of Digits in Base K](https://leetcode.com/problems/sum-of-digits-in-base-k/) | Compute the sum of digits of a number in base `k`. | Easy |
| 1844 | [Replace All Digits with Characters](https://leetcode.com/problems/replace-all-digits-with-characters/) | Replace digits in a string with shifted characters. | Easy |
| 1859 | [Sorting the Sentence](https://leetcode.com/problems/sorting-the-sentence/) | Sort a shuffled sentence based on word indices. | Easy |
| 1863 | [Sum of All Subset XOR Totals](https://leetcode.com/problems/sum-of-all-subset-xor-totals/) | Compute the sum of XOR totals for all possible subsets. | Easy |
| 1876 | [Substrings of Size Three with Distinct Characters](https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/) | Count substrings of length 3 with all distinct characters. | Easy |
| 1880 | [Check if Word Equals Summation of Two Words](https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/) | Check if the numerical value of a word equals the sum of two other words. | Easy |
| 1897 | [Redistribute Characters to Make All Strings Equal](https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/) | Check if characters can be redistributed to make all strings equal. | Easy |
| 1903 | [Largest Odd Number in String](https://leetcode.com/problems/largest-odd-number-in-string/) | Find the largest odd number that is a substring of a given string. | Easy |
| 1909 | [Remove One Element to Make the Array Strictly Increasing](https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/) | Check if an array can be made strictly increasing by removing one element. | Easy |
| 1913 | [Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/) | Find the maximum difference between the products of two pairs in an array. | Easy |
| 1920 | [Build Array from Permutation](https://leetcode.com/problems/build-array-from-permutation/) | Build an array where each element is the value at the index given by the input array. | Easy |
| 1925 | [Count Square Sum Triples](https://leetcode.com/problems/count-square-sum-triples/) | Count triples of integers that form a Pythagorean triple. | Easy |
| 1929 | [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) | Concatenate an array with itself. | Easy |
| 1935 | [Maximum Number of Words You Can Type](https://leetcode.com/problems/maximum-number-of-words-you-can-type/) | Count the number of words that can be typed using a broken keyboard. | Easy |
| 1941 | [Check if All Characters Have Equal Number of Occurrences](https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/) | Check if all characters in a string have the same frequency. | Easy |
| 1945 | [Sum of Digits of String After Convert](https://leetcode.com/problems/sum-of-digits-of-string-after-convert/) | Compute the sum of digits after converting a string and repeating `k` times. | Easy |
| 1952 | [Three Divisors](https://leetcode.com/problems/three-divisors/) | Check if a number has exactly three divisors. | Easy |
| 1957 | [Delete Characters to Make Fancy String](https://leetcode.com/problems/delete-characters-to-make-fancy-string/) | Remove characters to ensure no three consecutive characters are the same. | Easy |
| 1961 | [Check If String Is a Prefix of Array](https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/) | Check if a string is a prefix of the concatenation of an array of strings. | Easy |
| 1967 | [Number of Strings That Appear as Substrings in Word](https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/) | Count how many strings appear as substrings in a word. | Easy |
| 1974 | [Minimum Time to Type Word Using Special Typewriter](https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/) | Calculate the minimum time to type a word on a special typewriter. | Easy |
| 1979 | [Find Greatest Common Divisor of Array](https://leetcode.com/problems/find-greatest-common-divisor-of-array/) | Find the GCD of the minimum and maximum elements in an array. | Easy |
| 1984 | [Minimum Difference Between Highest and Lowest of K Scores](https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/) | Find the minimum difference between the highest and lowest scores in any `k` elements. | Easy |
| 1991 | [Find the Middle Index in Array](https://leetcode.com/problems/find-the-middle-index-in-array/) | Find the index where the sum of elements on the left equals the sum on the right. | Easy |
| 1995 | [Count Special Quadruplets](https://leetcode.com/problems/count-special-quadruplets/) | Count quadruplets in an array where the sum of the first three equals the fourth. | Easy |
| 2000 | [Reverse Prefix of Word](https://leetcode.com/problems/reverse-prefix-of-word/) | Reverse the prefix of a word up to the first occurrence of a given character. | Easy |
| 2006 | [Count Number of Pairs With Absolute Difference K](https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/) | Count pairs in an array with an absolute difference of `k`. | Easy |
| 2011 | [Final Value of Variable After Performing Operations](https://leetcode.com/problems/final-value-of-variable-after-performing-operations/) | Compute the final value of a variable after performing operations. | Easy |
| 2016 | [Maximum Difference Between Increasing Elements](https://leetcode.com/problems/maximum-difference-between-increasing-elements/) | Find the maximum difference between two elements where the second is larger. | Easy |
| 2022 | [Convert 1D Array Into 2D Array](https://leetcode.com/problems/convert-1d-array-into-2d-array/) | Convert a 1D array into a 2D array with given dimensions. | Easy |
| 2032 | [Two Out of Three](https://leetcode.com/problems/two-out-of-three/) | Find elements that appear in at least two of three arrays. | Easy |
| 2037 | [Minimum Number of Moves to Seat Everyone](https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/) | Calculate the minimum moves to seat students in seats. | Easy |
| 2042 | [Check if Numbers Are Ascending in a Sentence](https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/) | Check if numbers in a sentence are in strictly ascending order. | Easy |
| 2047 | [Number of Valid Words in a Sentence](https://leetcode.com/problems/number-of-valid-words-in-a-sentence/) | Count valid words in a sentence based on specific rules. | Easy |
| 2053 | [Kth Distinct String in an Array](https://leetcode.com/problems/kth-distinct-string-in-an-array/) | Find the `k`-th distinct string in an array. | Easy |
| 2057 | [Smallest Index With Equal Value](https://leetcode.com/problems/smallest-index-with-equal-value/) | Find the smallest index where the index equals the value. | Easy |
| 2062 | [Count Vowel Substrings of a String](https://leetcode.com/problems/count-vowel-substrings-of-a-string/) | Count substrings that contain all five vowels. | Easy |
| 2068 | [Check Whether Two Strings are Almost Equivalent](https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/) | Check if two strings have almost equivalent character frequencies. | Easy |
| 2073 | [Time Needed to Buy Tickets](https://leetcode.com/problems/time-needed-to-buy-tickets/) | Calculate the time needed for a person to buy tickets in a queue. | Easy |
| 2078 | [Two Furthest Houses With Different Colors](https://leetcode.com/problems/two-furthest-houses-with-different-colors/) | Find the maximum distance between two houses with different colors. | Easy |
| 2085 | [Count Common Words With One Occurrence](https://leetcode.com/problems/count-common-words-with-one-occurrence/) | Count words that appear exactly once in both of two arrays. | Easy |
| 2089 | [Find Target Indices After Sorting Array](https://leetcode.com/problems/find-target-indices-after-sorting-array/) | Find all indices of a target value in a sorted array. | Easy |
| 2094 | [Finding 3-Digit Even Numbers](https://leetcode.com/problems/finding-3-digit-even-numbers/) | Find all possible 3-digit even numbers from a given set of digits. | Easy |
| 2099 | [Find Subsequence of Length K With the Largest Sum](https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/) | Find a subsequence of length `k` with the largest sum. | Easy |
| 2103 | [Rings and Rods](https://leetcode.com/problems/rings-and-rods/) | Count rods that have all three colors of rings. | Easy |
| 2108 | [Find First Palindromic String in the Array](https://leetcode.com/problems/find-first-palindromic-string-in-the-array/) | Find the first palindromic string in an array. | Easy |
| 2114 | [Maximum Number of Words Found in Sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/) | Find the maximum number of words in any sentence. | Easy |
| 2119 | [A Number After a Double Reversal](https://leetcode.com/problems/a-number-after-a-double-reversal/) | Check if a number remains the same after reversing it twice. | Easy |
| 2124 | [Check if All A's Appears Before All B's](https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/) | Check if all 'A's appear before all 'B's in a string. | Easy |
| 2129 | [Capitalize the Title](https://leetcode.com/problems/capitalize-the-title/) | Capitalize words in a title according to specific rules. | Easy |
| 2133 | [Check if Every Row and Column Contains All Numbers](https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/) | Check if a matrix contains all numbers from 1 to `n` in every row and column. | Easy |
| 2138 | [Divide a String Into Groups of Size k](https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/) | Divide a string into groups of size `k`, filling the last group if needed. | Easy |
| 2144 | [Minimum Cost of Buying Candies With Discount](https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/) | Compute the minimum cost to buy candies with a "buy two get one free" discount. | Easy |
| 2148 | [Count Elements With Strictly Smaller and Greater Elements](https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/) | Count elements that have both strictly smaller and greater elements. | Easy |
| 2154 | [Keep Multiplying Found Values by Two](https://leetcode.com/problems/keep-multiplying-found-values-by-two/) | Keep multiplying a value by 2 as long as it exists in an array. | Easy |
| 2160 | [Minimum Sum of Four Digit Number After Splitting Digits](https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/) | Find the minimum sum by splitting a 4-digit number into two 2-digit numbers. | Easy |
| 2164 | [Sort Even and Odd Indices Independently](https://leetcode.com/problems/sort-even-and-odd-indices-independently/) | Sort elements at even and odd indices independently. | Easy |
| 2169 | [Count Operations to Obtain Zero](https://leetcode.com/problems/count-operations-to-obtain-zero/) | Count operations to make two numbers zero by subtracting the smaller from the larger. | Easy |
| 2176 | [Count Equal and Divisible Pairs in an Array](https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/) | Count pairs where elements are equal and their indices' product is divisible by `k`. | Easy |
| 2180 | [Count Integers With Even Digit Sum](https://leetcode.com/problems/count-integers-with-even-digit-sum/) | Count integers up to `n` with an even digit sum. | Easy |
| 2185 | [Counting Words With a Given Prefix](https://leetcode.com/problems/counting-words-with-a-given-prefix/) | Count words that start with a given prefix. | Easy |
| 2190 | [Most Frequent Number Following Key In an Array](https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/) | Find the most frequent number that follows a key in an array. | Easy |
| 2194 | [Cells in a Range on an Excel Sheet](https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/) | List all cells in a given range on an Excel sheet. | Easy |
| 2200 | [Find All K-Distant Indices in an Array](https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/) | Find all indices that are within `k` distance of an index containing a key. | Easy |
| 2206 | [Divide Array Into Equal Pairs](https://leetcode.com/problems/divide-array-into-equal-pairs/) | Check if an array can be divided into pairs of equal elements. | Easy |
| 2210 | [Count Hills and Valleys in an Array](https://leetcode.com/problems/count-hills-and-valleys-in-an-array/) | Count hills and valleys in an array based on adjacent elements. | Easy |
| 2215 | [Find the Difference of Two Arrays](https://leetcode.com/problems/find-the-difference-of-two-arrays/) | Find elements in one array that are not in the other. | Easy |
| 2220 | [Minimum Bit Flips to Convert Number](https://leetcode.com/problems/minimum-bit-flips-to-convert-number/) | Count the minimum bit flips to convert one number to another. | Easy |
| 2224 | [Minimum Number of Operations to Convert Time](https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/) | Compute the minimum operations to convert one time to another. | Easy |
| 2231 | [Largest Number After Digit Swaps by Parity](https://leetcode.com/problems/largest-number-after-digit-swaps-by-parity/) | Find the largest number by swapping digits of the same parity. | Easy |
| 2235 | [Add Two Integers](https://leetcode.com/problems/add-two-integers/) | Add two integers without using the `+` or `-` operators. | Easy |
| 2236 | [Root Equals Sum of Children](https://leetcode.com/problems/root-equals-sum-of-children/) | Check if a binary tree's root value equals the sum of its children's values. | Easy |
| 2239 | [Find Closest Number to Zero](https://leetcode.com/problems/find-closest-number-to-zero/) | Find the number closest to zero in an array. | Easy |
| 2243 | [Calculate Digit Sum of a String](https://leetcode.com/problems/calculate-digit-sum-of-a-string/) | Calculate the digit sum of a string by grouping and summing digits. | Easy |
| 2248 | [Intersection of Multiple Arrays](https://leetcode.com/problems/intersection-of-multiple-arrays/) | Find the intersection of multiple arrays. | Easy |
| 2255 | [Count Prefixes of a Given String](https://leetcode.com/problems/count-prefixes-of-a-given-string/) | Count how many strings are prefixes of a given string. | Easy |
| 2259 | [Remove Digit From Number to Maximize Result](https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/) | Remove one occurrence of a digit to maximize the resulting number. | Easy |
| 2264 | [Largest 3-Same-Digit Number in String](https://leetcode.com/problems/largest-3-same-digit-number-in-string/) | Find the largest number formed by three consecutive same digits. | Easy |
| 2269 | [Find the K-Beauty of a Number](https://leetcode.com/problems/find-the-k-beauty-of-a-number/) | Count substrings of a number that are divisors of the number. | Easy |
| 2278 | [Percentage of Letter in String](https://leetcode.com/problems/percentage-of-letter-in-string/) | Compute the percentage of occurrences of a letter in a string. | Easy |
| 2283 | [Check if Number Has Equal Digit Count and Digit Value](https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/) | Check if the count of each digit equals the digit's value. | Easy |
| 2287 | [Rearrange Characters to Make Target String](https://leetcode.com/problems/rearrange-characters-to-make-target-string/) | Find the maximum number of times a target string can be formed. | Easy |
| 2293 | [Min Max Game](https://leetcode.com/problems/min-max-game/) | Simulate a game to find the final number in an array. | Easy |
| 2299 | [Strong Password Checker II](https://leetcode.com/problems/strong-password-checker-ii/) | Check if a password meets specific strength criteria. | Easy |
| 2303 | [Calculate Amount Paid in Taxes](https://leetcode.com/problems/calculate-amount-paid-in-taxes/) | Calculate the total tax paid given income and tax brackets. | Easy |
| 2309 | [Greatest English Letter in Upper and Lower Case](https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/) | Find the greatest letter that appears in both upper and lower case. | Easy |
| 2315 | [Count Asterisks](https://leetcode.com/problems/count-asterisks/) | Count asterisks in a string that are not between pairs of vertical bars. | Easy |
| 2319 | [Check if Matrix Is X-Matrix](https://leetcode.com/problems/check-if-matrix-is-x-matrix/) | Check if a matrix is an X-matrix (non-zero on diagonals, zero elsewhere). | Easy |
| 2325 | [Decode the Message](https://leetcode.com/problems/decode-the-message/) | Decode a message using a given key to map characters. | Easy |
| 2331 | [Evaluate Boolean Binary Tree](https://leetcode.com/problems/evaluate-boolean-binary-tree/) | Evaluate a boolean binary tree with AND, OR, and leaf nodes. | Easy |
| 2335 | [Minimum Amount of Time to Fill Cups](https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/) | Calculate the minimum time to fill three cups with given amounts. | Easy |
| 2341 | [Maximum Number of Pairs in Array](https://leetcode.com/problems/maximum-number-of-pairs-in-array/) | Find the maximum number of pairs that can be formed from an array. | Easy |
| 2347 | [Best Poker Hand](https://leetcode.com/problems/best-poker-hand/) | Determine the best poker hand from a given set of cards. | Easy |
| 2351 | [First Letter to Appear Twice](https://leetcode.com/problems/first-letter-to-appear-twice/) | Find the first letter that appears twice in a string. | Easy |
| 2357 | [Make Array Zero by Subtracting Equal Amounts](https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/) | Find the minimum operations to make all elements zero by subtracting equal amounts. | Easy |
| 2363 | [Merge Similar Items](https://leetcode.com/problems/merge-similar-items/) | Merge two lists of items by summing values for the same IDs. | Easy |
| 2367 | [Number of Arithmetic Triplets](https://leetcode.com/problems/number-of-arithmetic-triplets/) | Count triplets in an array that form an arithmetic sequence. | Easy |
| 2373 | [Largest Local Values in a Matrix](https://leetcode.com/problems/largest-local-values-in-a-matrix/) | Find the largest value in each 3x3 submatrix of a matrix. | Easy |
| 2379 | [Minimum Recolors to Get K Consecutive Black Blocks](https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/) | Find the minimum recolors to get `k` consecutive black blocks. | Easy |
| 2383 | [Minimum Hours of Training to Win a Competition](https://leetcode.com/problems/minimum-hours-of-training-to-win-a-competition/) | Calculate the minimum hours of training to win a competition. | Easy |
| 2389 | [Longest Subsequence With Limited Sum](https://leetcode.com/problems/longest-subsequence-with-limited-sum/) | Find the longest subsequence with a sum less than or equal to a given value. | Easy |
| 2395 | [Find Subarrays With Equal Sum](https://leetcode.com/problems/find-subarrays-with-equal-sum/) | Check if there are two subarrays with equal sums. | Easy |
| 2399 | [Check Distances Between Same Letters](https://leetcode.com/problems/check-distances-between-same-letters/) | Check if the distances between same letters match a given array. | Easy |
| 2404 | [Most Frequent Even Element](https://leetcode.com/problems/most-frequent-even-element/) | Find the most frequent even number in an array. | Easy |
| 2413 | [Smallest Even Multiple](https://leetcode.com/problems/smallest-even-multiple/) | Find the smallest number that is a multiple of both 2 and `n`. | Easy |
| 2418 | [Sort the People](https://leetcode.com/problems/sort-the-people/) | Sort people by height in descending order. | Easy |
| 2423 | [Remove Letter To Equalize Frequency](https://leetcode.com/problems/remove-letter-to-equalize-frequency/) | Check if removing one letter can make all character frequencies equal. | Easy |
| 2427 | [Number of Common Factors](https://leetcode.com/problems/number-of-common-factors/) | Count the number of common factors between two integers. | Easy |
| 2432 | [The Employee That Worked on the Longest Task](https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/) | Find the employee who worked on the longest task. | Easy |
| 2437 | [Number of Valid Clock Times](https://leetcode.com/problems/number-of-valid-clock-times/) | Count valid clock times given a string with `?` placeholders. | Easy |
| 2441 | [Largest Positive Integer That Exists With Its Negative](https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/) | Find the largest positive integer that also exists as its negative in an array. | Easy |
| 2446 | [Determine if Two Events Have Conflict](https://leetcode.com/problems/determine-if-two-events-have-conflict/) | Check if two events overlap in time. | Easy |
| 2451 | [Odd String Difference](https://leetcode.com/problems/odd-string-difference/) | Find the string with a different difference array compared to others. | Easy |
| 2455 | [Average Value of Even Numbers That Are Divisible by Three](https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/) | Compute the average of even numbers divisible by three. | Easy |
| 2460 | [Apply Operations to an Array](https://leetcode.com/problems/apply-operations-to-an-array/) | Apply operations to double adjacent equal elements and move zeros to the end. | Easy |
| 2465 | [Number of Distinct Averages](https://leetcode.com/problems/number-of-distinct-averages/) | Count distinct averages of pairs from an array. | Easy |
| 2469 | [Convert the Temperature](https://leetcode.com/problems/convert-the-temperature/) | Convert a temperature from Celsius to Kelvin and Fahrenheit. | Easy |
| 2475 | [Number of Unequal Triplets](https://leetcode.com/problems/number-of-unequal-triplets/) | Count triplets with all distinct elements. | Easy |
| 2481 | [Minimum Cuts to Divide a Circle](https://leetcode.com/problems/minimum-cuts-to-divide-a-circle/) | Find the minimum cuts to divide a circle into `n` equal parts. | Easy |
| 2485 | [Find the Pivot Integer](https://leetcode.com/problems/find-the-pivot-integer/) | Find an integer where the sum of numbers before and after it is equal. | Easy |
| 2490 | [Circular Sentence](https://leetcode.com/problems/circular-sentence/) | Check if a sentence is circular (last letter of each word matches the first letter of the next). | Easy |
| 2496 | [Maximum Value of a String in an Array](https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/) | Find the maximum value of strings (length or numeric value). | Easy |
| 2500 | [Delete Greatest Value in Each Row](https://leetcode.com/problems/delete-greatest-value-in-each-row/) | Compute the sum of the maximum values deleted from each row of a matrix. | Easy |
| 2506 | [Count Pairs Of Similar Strings](https://leetcode.com/problems/count-pairs-of-similar-strings/) | Count pairs of strings with the same set of characters. | Easy |
| 2511 | [Maximum Enemy Forts That Can Be Captured](https://leetcode.com/problems/maximum-enemy-forts-that-can-be-captured/) | Find the maximum number of enemy forts that can be captured. | Easy |
| 2515 | [Shortest Distance to Target String in a Circular Array](https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/) | Find the shortest distance to a target string in a circular array. | Easy |
| 2520 | [Count the Digits That Divide a Number](https://leetcode.com/problems/count-the-digits-that-divide-a-number/) | Count digits in a number that divide the number evenly. | Easy |
| 2525 | [Categorize Box According to Criteria](https://leetcode.com/problems/categorize-box-according-to-criteria/) | Categorize a box as bulky, heavy, both, or neither based on dimensions and mass. | Easy |
| 2529 | [Maximum Count of Positive Integer and Negative Integer](https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/) | Find the maximum count of positive or negative integers in an array. | Easy |
| 2535 | [Difference Between Element Sum and Digit Sum of an Array](https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/) | Compute the absolute difference between the sum of elements and the sum of their digits. | Easy |
| 2540 | [Minimum Common Value](https://leetcode.com/problems/minimum-common-value/) | Find the minimum common element in two sorted arrays. | Easy |
| 2544 | [Alternating Digit Sum](https://leetcode.com/problems/alternating-digit-sum/) | Compute the alternating sum of digits in a number. | Easy |
| 2549 | [Count Distinct Numbers on Board](https://leetcode.com/problems/count-distinct-numbers-on-board/) | Count distinct numbers on a board after applying specific rules. | Easy |
| 2553 | [Separate the Digits in an Array](https://leetcode.com/problems/separate-the-digits-in-an-array/) | Separate the digits of each number in an array. | Easy |
| 2558 | [Take Gifts From the Richest Pile](https://leetcode.com/problems/take-gifts-from-the-richest-pile/) | Simulate taking gifts from the richest pile by reducing its square root. | Easy |
| 2562 | [Find the Array Concatenation Value](https://leetcode.com/problems/find-the-array-concatenation-value/) | Compute the concatenation value of an array by pairing elements. | Easy |
| 2566 | [Maximum Difference by Remapping a Digit](https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/) | Find the maximum difference by remapping one digit in a number. | Easy |
| 2570 | [Merge Two 2D Arrays by Summing Values](https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/) | Merge two 2D arrays by summing values for the same IDs. | Easy |
| 2574 | [Left and Right Sum Differences](https://leetcode.com/problems/left-and-right-sum-differences/) | Compute the absolute difference between left and right sums for each index. | Easy |
| 2578 | [Split With Minimum Sum](https://leetcode.com/problems/split-with-minimum-sum/) | Split a number into two numbers with the minimum possible sum. | Easy |
| 2582 | [Pass the Pillow](https://leetcode.com/problems/pass-the-pillow/) | Determine who holds a pillow after passing it for a given time. | Easy |
| 2586 | [Count the Number of Vowel Strings in Range](https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/) | Count strings in a range that start and end with a vowel. | Easy |
| 2591 | [Distribute Money to Maximum Children](https://leetcode.com/problems/distribute-money-to-maximum-children/) | Maximize the number of children who receive exactly 8 dollars. | Easy |
| 2595 | [Number of Even and Odd Bits](https://leetcode.com/problems/number-of-even-and-odd-bits/) | Count even and odd-indexed 1 bits in a number's binary representation. | Easy |
| 2600 | [K Items With the Maximum Sum](https://leetcode.com/problems/k-items-with-the-maximum-sum/) | Find the maximum sum by selecting `k` items from three types. | Easy |
| 2605 | [Form Smallest Number From Two Digit Arrays](https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/) | Form the smallest number using digits from two arrays. | Easy |
| 2609 | [Find the Longest Balanced Substring of a Binary String](https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/) | Find the longest substring with an equal number of 0s and 1s. | Easy |
| 2614 | [Prime In Diagonal](https://leetcode.com/problems/prime-in-diagonal/) | Find the largest prime number on the diagonals of a matrix. | Easy |
| 2618 | [Check if Object Instance of Class](https://leetcode.com/problems/check-if-object-instance-of-class/) | Check if an object is an instance of a given class or its subclasses. | Easy |
| 2620 | [Counter](https://leetcode.com/problems/counter/) | Create a counter function that increments from a given initial value. | Easy |
| 2621 | [Sleep](https://leetcode.com/problems/sleep/) | Implement an async sleep function that pauses for a given number of milliseconds. | Easy |
| 2622 | [Cache With Time Limit](https://leetcode.com/problems/cache-with-time-limit/) | Implement a cache with a time limit for key-value pairs. | Easy |
| 2623 | [Memoize](https://leetcode.com/problems/memoize/) | Implement a memoization function to cache function results. | Easy |
| 2626 | [Array Reduce Transformation](https://leetcode.com/problems/array-reduce-transformation/) | Implement a reduce function to transform an array with a given reducer function. | Easy |
| 2628 | [JSON Deep Equal](https://leetcode.com/problems/json-deep-equal/) | Check if two JSON objects are deeply equal. | Easy |
| 2629 | [Function Composition](https://leetcode.com/problems/function-composition/) | Implement function composition to apply functions in sequence. | Easy |
| 2631 | [Group By](https://leetcode.com/problems/group-by/) | Implement a groupBy function to group array elements by a key function. | Easy |
| 2632 | [Curry](https://leetcode.com/problems/curry/) | Implement a curry function that transforms a function to accept arguments one at a time. | Easy |
| 2634 | [Filter Elements from Array](https://leetcode.com/problems/filter-elements-from-array/) | Implement a filter function to return elements from an array that satisfy a condition. | Easy |
| 2635 | [Apply Transform Over Each Element in Array](https://leetcode.com/problems/apply-transform-over-each-element-in-array/) | Apply a given transformation function to each element in an array. | Easy |
| 2637 | [Promise Time Limit](https://leetcode.com/problems/promise-time-limit/) | Implement a function that limits the execution time of a promise, rejecting if it exceeds the limit. | Easy |
| 2648 | [Generate Fibonacci Sequence](https://leetcode.com/problems/generate-fibonacci-sequence/) | Implement a generator function to yield the Fibonacci sequence. | Easy |
| 2649 | [Nested Array Generator](https://leetcode.com/problems/nested-array-generator/) | Implement a generator to flatten a nested array. | Easy |
| 2665 | [Counter II](https://leetcode.com/problems/counter-ii/) | Implement a counter object with increment, decrement, and reset methods. | Easy |
| 2666 | [Allow One Function Call](https://leetcode.com/problems/allow-one-function-call/) | Modify a function to be called only once, returning undefined on subsequent calls. | Easy |
| 2667 | [Create Hello World Function](https://leetcode.com/problems/create-hello-world-function/) | Create a function that returns "Hello World". | Easy |
| 2675 | [Array Prototype Last](https://leetcode.com/problems/array-prototype-last/) | Add a method to Array.prototype to return the last element or -1 if empty. | Easy |
| 2677 | [Chunk Array](https://leetcode.com/problems/chunk-array/) | Implement a function to split an array into chunks of a given size. | Easy |
| 2693 | [Call Function with Custom Context](https://leetcode.com/problems/call-function-with-custom-context/) | Implement a function to call another function with a specified context. | Easy |
| 2694 | [Event Emitter](https://leetcode.com/problems/event-emitter/) | Implement an event emitter with subscribe and emit functionality. | Easy |
| 2695 | [Array Wrapper](https://leetcode.com/problems/array-wrapper/) | Create a wrapper class for an array with custom string and value conversions. | Easy |
| 2703 | [Return Length of Arguments Passed](https://leetcode.com/problems/return-length-of-arguments-passed/) | Return the number of arguments passed to a function. | Easy |
| 2704 | [To Be Or Not To Be](https://leetcode.com/problems/to-be-or-not-to-be/) | Implement a function that checks if a value equals a given value. | Easy |
| 2705 | [Compact Object](https://leetcode.com/problems/compact-object/) | Remove falsy values from an object or array recursively. | Easy |
| 2715 | [Timeout Cancellation](https://leetcode.com/problems/timeout-cancellation/) | Implement a function that executes a callback after a delay with cancellation. | Easy |
| 2721 | [Execute Asynchronous Functions in Parallel](https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/) | Execute an array of asynchronous functions in parallel and collect results. | Easy |
| 2722 | [Join Two Arrays by ID](https://leetcode.com/problems/join-two-arrays-by-id/) | Merge two arrays of objects by matching IDs and combining properties. | Easy |
| 2723 | [Add Two Promises](https://leetcode.com/problems/add-two-promises/) | Add the resolved values of two promises. | Easy |
| 2724 | [Sort By](https://leetcode.com/problems/sort-by/) | Sort an array based on a provided comparison function. | Easy |
| 2725 | [Interval Cancellation](https://leetcode.com/problems/interval-cancellation/) | Implement a function that repeatedly calls a callback at intervals with cancellation. | Easy |
| 2726 | [Calculator with Method Chaining](https://leetcode.com/problems/calculator-with-method-chaining/) | Implement a calculator class that supports method chaining for arithmetic operations. | Easy |
| 2727 | [Is Object Empty](https://leetcode.com/problems/is-object-empty/) | Check if an object or array is empty. | Easy |
| 22   | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) | Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses. | Medium |
| 8    | [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) | Implement the `myAtoi(string s)` function, which converts a string to a 32-bit signed integer with proper handling of whitespace, sign, non-digit characters, and overflow. | Medium |
| 7    | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | Reverse digits of a signed 32-bit integer, returning 0 if the result overflows the signed 32-bit range. | Medium |
| 423	| [Reconstruct Original Digits from English] (https://leetcode.com/problems/reconstruct-original-digits-from-english/description/) |	Given a string s containing an out-of-order English representation of digits 0–9, return the digits in ascending order. | Medium |
| 482    | [License Key Formatting]  (https://leetcode.com/problems/license-key-formatting/description/) | Given a license key string containing letters, digits, and dashes, and an integer k, reformat the string so every group (except possibly the first) has exactly k uppercase alphanumeric characters, separated by dashes, and remove extra dashes. | Easy |
| 3542 | [Minimum Operations to Convert All Elements to Zero] (https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/description/) | Given an array nums of size n consisting of non-negative integers, the task is to apply some (possibly zero) operations on the array so that all elements become zero. In one operation, you can select a subarray [i, j] (where 0 <= i <= j < n) and set all occurrences of the minimum non-negative integer in that subarray to zero. Return the minimum number of operations required to make all elements in the array zero. | Medium |
| 1186 | [Maximum Subarray Sum with One Deletion] (https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/description/) | Given an array of integers, return the maximum sum for a non-empty subarray (contiguous elements) with at most one element deletion. In other words, choose a subarray and optionally delete one element so that there is still at least one element left and the sum of the remaining elements is maximum possible. The subarray needs to be non-empty after deletion. | Medium |
| 954 | [Array of Doubled Pairs] (https://leetcode.com/problems/array-of-doubled-pairs/description/) | Given an integer array of even length arr, return true if it is possible to reorder arr such that arr[2 * i + 1] = 2 * arr[2 * i] for every 0 <= i < len(arr) / 2, or false otherwise. | Medium |
| 2727 | [Vowels Game in a String](https://leetcode.com/problems/vowels-game-in-a-string/) | Alice and Bob take turns removing substrings: Alice removes those with odd vowels, Bob with even vowels. The first player who can't make a move loses. Return true if Alice wins, assuming both play optimally. The game is played on the English vowels "a", "e", "i", "o", and "u". | Medium  
| 997 | [Find the Town Judge](https://leetcode.com/problems/find-the-town-judge/) | In a town of n people, the judge trusts nobody, and everybody trusts the judge. Return the label of the town judge if exists, otherwise -1. | Easy  
| 1971 | [Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/) | Check if there is a valid path from vertex source to vertex destination in a bidirectional graph with n vertices. | Easy  
| 3456 | [Find Special Substring of Length K](https://leetcode.com/problems/find-special-substring-of-length-k/) | Determine if there exists a substring of length exactly k in s that consists of only one distinct character, and is bordered by other characters. | Easy  
| 1736 | [Latest Time by Replacing Hidden Digits](https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/) | Given a time string hh:mm with some hidden digits marked by '?', return the latest valid time you can get by replacing the hidden digits. | Easy  
| 456 | [132 Pattern](https://leetcode.com/problems/132-pattern/) | Check if the array contains a 132 pattern: indices i < j < k with nums[i] < nums[k] < nums[j]. | Medium  
| 1358 | [Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/) | Count the number of substrings of s (consisting only of 'a', 'b', 'c') that contain at least one occurrence of all three characters. | Medium  
| 775 | [Global and Local Inversions](https://leetcode.com/problems/global-and-local-inversions/) | Check if the number of global inversions in the permutation array nums (where global inversion is any i < j with nums[i] > nums[j]) equals the number of local inversions (i with nums[i] > nums[i+1]). | Medium  
| 832 | [Flipping an Image](https://leetcode.com/problems/flipping-an-image/) | Flip the binary image horizontally and invert each pixel. | Easy  
| 1512 | [Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/) | Count the number of good pairs (i, j) with i < j and nums[i] == nums[j]. | Easy  
| 3484 | [Design Spreadsheet](https://leetcode.com/problems/design-spreadsheet/) | Implement a spreadsheet class that supports setting, resetting cell values, and evaluating simple formulas like "=A1+B2". Unset cell values are 0. | Medium  
| 2798 | [Number of Employees Who Met the Target](https://leetcode.com/problems/number-of-employees-who-met-the-target/) | Return the number of employees who have worked at least target hours. | Easy  
| 2761 | [Prime Pairs With Target Sum](https://leetcode.com/problems/prime-pairs-with-target-sum/) | Find all pairs of prime numbers x ≤ y such that x + y = n and both x and y are primes. | Medium  
| 1629 | [Slowest Key](https://leetcode.com/problems/slowest-key/) | Find the key pressed with the longest duration. If there is a tie, return the lexicographically larger key. | Easy  
| 1491 | [Average Salary Excluding the Minimum and Maximum Salary](https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/) | Compute the average salary of employees excluding the minimum and maximum salary. | Easy  
| 2828 | [Check if a String Is an Acronym of Words](https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/) | See if string s is an acronym of the words array, formed by the first character of each word in order. | Easy
| 654  | [Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/description/) | Construct the maximum binary tree from an array by recursively selecting the maximum element as the root and building left and right subtrees from subarrays split by the maximum element. | Medium   |
| 144  | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/description/) | Return the preorder traversal (root, left, right) of a binary tree's nodes' values. | Easy     |
| 94   | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/description/) | Return the inorder traversal (left, root, right) of a binary tree's nodes' values. | Easy     |
| 965  | [Univalued Binary Tree](https://leetcode.com/problems/univalued-binary-tree/description/) | Determine whether all nodes in a binary tree have the same value. | Easy     |
| 563  | [Binary Tree Tilt](https://leetcode.com/problems/binary-tree-tilt/description/) | Calculate the sum of all nodes' tilts, where the tilt of a node is the absolute difference between the sum of values of its left and right subtrees. | Easy     |
| 101  | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/description/) | Check if a binary tree is a mirror of itself (symmetric around its center). | Easy     |
| 100  | [Same Tree](https://leetcode.com/problems/same-tree/description/) | Check if two binary trees are structurally identical with the same node values. | Easy     |
| 3131 | [Find the Integer Added to Array I](https://leetcode.com/problems/find-the-integer-added-to-array-i/description/) | Given two arrays where one is formed by adding a fixed integer to all elements of the other, find that integer. | Easy     |
| 345  | [Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/description/) | Reverse only the vowels in a given string, preserving the position of other characters. | Easy     |
| 2148 | [Count Elements With Strictly Smaller and Greater Elements](https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/description/) | Count elements in an array that have both a strictly smaller and strictly greater element in the array. | Easy     |
| 2239 | [Find Closest Number to Zero](https://leetcode.com/problems/find-closest-number-to-zero/description/) | Find the number closest to zero in an array; if there are ties, return the larger number. | Easy     |
| 2365 | [Task Scheduler II](https://leetcode.com/problems/task-scheduler-ii/description/) | Calculate the minimum number of days to complete tasks in order with a cooldown period between same task types. | Medium   |
| 2079 | [Watering Plants](https://leetcode.com/problems/watering-plants/description/) | Determine the number of steps required to water plants in order with a limited watering can capacity. | Medium   |
| 2515 | [Shortest Distance to Target String in a Circular Array](https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/description/) | Find shortest distance to reach target string in a circular array from a start index. | Easy     |
| 257  | [Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/description/) | Return all root-to-leaf paths in a binary tree. | Easy     |
| 226  | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/) | Invert a binary tree, swapping left and right children of all nodes. | Easy     |
| 110  | [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/description/) | Check if a binary tree is height-balanced, i.e., difference in heights of left and right subtrees of every node is at most 1. | Easy     |
| 617  | [Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/description/) | Merge two binary trees by summing overlapping nodes and reusing non-null nodes. | Easy     |
| 2900 | [Longest Unequal Adjacent Groups Subsequence I](https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/description/) | Find the longest subsequence from words such that adjacent strings differ in associated binary groups. | Easy     |
| 1668 | [Maximum Repeating Substring](https://leetcode.com/problems/maximum-repeating-substring/description/) | Find the maximum k such that word repeated k times is a substring of sequence. | Easy     |
| 1137 | [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/description/) | Compute the n-th Tribonacci number where each term is the sum of the preceding three terms. | Easy     |
| 1025 | [Divisor Game](https://leetcode.com/problems/divisor-game/description/) | Determine if the first player wins a game where players subtract divisors of n. | Easy     |
| 746  | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/) | Find minimum cost to reach the top of stairs with steps of one or two stairs at a time. | Easy     |
| 509  | [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/description/) | Compute the n-th Fibonacci number. | Easy     |
| 338  | [Counting Bits](https://leetcode.com/problems/counting-bits/description/) | For each number up to n, count the number of 1 bits in binary representation. | Easy     |
| 2855 | [Minimum Right Shifts to Sort the Array](https://leetcode.com/problems/minimum-right-shifts-to-sort-the-array/description/) | Find minimum number of right shifts to sort a distinct integer array or return -1 if impossible. | Easy     |
| 3043 | [Find the Length of the Longest Common Prefix](https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/description/) | Find longest common prefix length among all pairs from two integer arrays. | Medium   |
| 3158 | [Find the XOR of Numbers Which Appear Twice](https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/description/) | Return XOR of numbers appearing exactly twice in the array. | Easy     |
| 2154 | [Keep Multiplying Found Values by Two](https://leetcode.com/problems/keep-multiplying-found-values-by-two/description/) | Repeatedly multiply a found value by two as long as it is found in the array; return final value. | Easy     |
| 682  | [Baseball Game](https://leetcode.com/problems/baseball-game/description/) | Calculate total score given operations to record, double, add, or remove previous scores. | Easy     |
| 3227 | [Vowels Game in a String](https://leetcode.com/problems/vowels-game-in-a-string/description/) | Determine the winner of a game removing substrings with counts of vowels varying by turn. | Medium   |
| 997  | [Find the Town Judge](https://leetcode.com/problems/find-the-town-judge/description/) | Identify the town judge who trusts no one but is trusted by everyone else. | Easy     |
| 1971 | [Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/description/) | Check if a valid path exists between two vertices in a bi-directional graph. | Easy     |
| 3456 | [Find Special Substring of Length K](https://leetcode.com/problems/find-special-substring-of-length-k/description/) | Determine if a substring of length k exists with one distinct character and differing neighbors. | Easy     |
| 1736 | [Latest Time by Replacing Hidden Digits](https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/description/) | Replace hidden digits in a time string to get the latest valid time. | Easy     |
| 456  | [132 Pattern](https://leetcode.com/problems/132-pattern/description/) | Check if there exists a subsequence of three numbers following the 132 pattern. | Medium   |
| 1358 | [Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/) | Count substrings containing at least one of each character 'a', 'b', and 'c'. | Medium   |
| 775  | [Global and Local Inversions](https://leetcode.com/problems/global-and-local-inversions/description/) | Verify if the number of global inversions equals the number of local inversions in an array. | Medium   |
| 832  | [Flipping an Image](https://leetcode.com/problems/flipping-an-image/description/) | Flip a binary matrix horizontally and invert values (0 ↔ 1). | Easy     |
| 1512 | [Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/description/) | Count pairs (i, j) with i < j and nums[i] == nums[j]. | Easy     |
| 3484 | [Design Spreadsheet](https://leetcode.com/problems/design-spreadsheet/description/) | Implement spreadsheet class with set, reset, and formula evaluation capabilities. | Medium   |
| 2798 | [Number of Employees Who Met the Target](https://leetcode.com/problems/number-of-employees-who-met-the-target/description/) | Return count of employees who met or exceeded given target hours. | Easy     |
| 2761 | [Prime Pairs With Target Sum](https://leetcode.com/problems/prime-pairs-with-target-sum/description/) | Find all pairs of primes x ≤ y where x + y = n. | Medium   |
| 1629 | [Slowest Key](https://leetcode.com/problems/slowest-key/description/) | Return the key pressed for the longest duration or lexicographically largest in case of ties. | Easy     |
| 1491 | [Average Salary Excluding the Minimum and Maximum Salary](https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/description/) | Compute average salary excluding the minimum and maximum. | Easy     |
| 2828 | [Check if a String Is an Acronym of Words](https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/description/) | Verify if string s is an acronym formed from the first characters of words array. | Easy     |



## Usage

To use these solutions:

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Run the Code**:
   - Ensure you have Go installed (version 1.16 or later recommended).
   - The `main.go` file contains all solutions. You can test individual functions by calling them in the `main` function or writing test cases.
   - Run the program using:
     ```bash
     go run main.go
     ```

3. **Testing**:
   - Each function is designed to match LeetCode's input/output specifications.
   - You can create test cases using Go's `testing` package or manually call functions with sample inputs.

## Notes

- **Code Organization**: All solutions are currently in a single `main.go` file for simplicity. For larger projects, consider organizing functions into separate files or packages by problem category (e.g., arrays, strings, trees).
- **Performance**: Solutions aim for clarity and correctness. Some may not be fully optimized for time or space complexity but serve as a starting point for further improvement.
- **Contributions**: Feel free to contribute by adding new solutions, optimizing existing ones, or improving documentation. Please follow the contribution guidelines (to be added).

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
<!---LeetCode Topics Start-->
# LeetCode Topics
## Array
|  |
| ------- |
| [0917-boats-to-save-people](https://github.com/Ivan-Lapin/LeetCode/tree/master/0917-boats-to-save-people) |
| [1781-check-if-two-string-arrays-are-equivalent](https://github.com/Ivan-Lapin/LeetCode/tree/master/1781-check-if-two-string-arrays-are-equivalent) |
| [1830-count-good-meals](https://github.com/Ivan-Lapin/LeetCode/tree/master/1830-count-good-meals) |
| [2190-count-common-words-with-one-occurrence](https://github.com/Ivan-Lapin/LeetCode/tree/master/2190-count-common-words-with-one-occurrence) |
| [2283-sort-even-and-odd-indices-independently](https://github.com/Ivan-Lapin/LeetCode/tree/master/2283-sort-even-and-odd-indices-independently) |
| [2624-difference-between-element-sum-and-digit-sum-of-an-array](https://github.com/Ivan-Lapin/LeetCode/tree/master/2624-difference-between-element-sum-and-digit-sum-of-an-array) |
| [2836-neither-minimum-nor-maximum](https://github.com/Ivan-Lapin/LeetCode/tree/master/2836-neither-minimum-nor-maximum) |
## String
|  |
| ------- |
| [1781-check-if-two-string-arrays-are-equivalent](https://github.com/Ivan-Lapin/LeetCode/tree/master/1781-check-if-two-string-arrays-are-equivalent) |
| [2190-count-common-words-with-one-occurrence](https://github.com/Ivan-Lapin/LeetCode/tree/master/2190-count-common-words-with-one-occurrence) |
## Sorting
|  |
| ------- |
| [0917-boats-to-save-people](https://github.com/Ivan-Lapin/LeetCode/tree/master/0917-boats-to-save-people) |
| [2283-sort-even-and-odd-indices-independently](https://github.com/Ivan-Lapin/LeetCode/tree/master/2283-sort-even-and-odd-indices-independently) |
| [2836-neither-minimum-nor-maximum](https://github.com/Ivan-Lapin/LeetCode/tree/master/2836-neither-minimum-nor-maximum) |
## Hash Table
|  |
| ------- |
| [1830-count-good-meals](https://github.com/Ivan-Lapin/LeetCode/tree/master/1830-count-good-meals) |
| [2190-count-common-words-with-one-occurrence](https://github.com/Ivan-Lapin/LeetCode/tree/master/2190-count-common-words-with-one-occurrence) |
## Counting
|  |
| ------- |
| [2190-count-common-words-with-one-occurrence](https://github.com/Ivan-Lapin/LeetCode/tree/master/2190-count-common-words-with-one-occurrence) |
## Math
|  |
| ------- |
| [2624-difference-between-element-sum-and-digit-sum-of-an-array](https://github.com/Ivan-Lapin/LeetCode/tree/master/2624-difference-between-element-sum-and-digit-sum-of-an-array) |
| [2630-alternating-digit-sum](https://github.com/Ivan-Lapin/LeetCode/tree/master/2630-alternating-digit-sum) |
## Two Pointers
|  |
| ------- |
| [0917-boats-to-save-people](https://github.com/Ivan-Lapin/LeetCode/tree/master/0917-boats-to-save-people) |
## Greedy
|  |
| ------- |
| [0917-boats-to-save-people](https://github.com/Ivan-Lapin/LeetCode/tree/master/0917-boats-to-save-people) |
<!---LeetCode Topics End-->