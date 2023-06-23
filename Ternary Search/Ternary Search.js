//normal
const TernarySearch = (left, right, target, arr) => {
  while (right >= left) {
    let mid1 = left + Math.floor((right - left) / 3)
    let mid2 = right - Math.floor((right - left) / 3)

    if (arr[mid1] == target) {
      return mid1
    }
    if (arr[mid2] == target) {
      return mid2
    }

    if (target < arr[mid1]) {
      right = mid1 - 1
    } else if (target > arr[mid2]) {
      left = mid2 + 1
    } else {
      left = mid1 + 1
      right = mid2 - 1
    }
  }

  return -1
}

let arr1 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
const left1 = 0
const right1 = arr1.length - 1
const target1 = 1
const result1 = TernarySearch(left1, right1, target1, arr1)

//recursive :
const ternarySearch = (left, right, target, arr) => {
  if (right >= left) {
    let mid1 = left + Math.floor((right - left) / 3)
    let mid2 = right - Math.floor((right - left) / 3)

    if (arr[mid1] == target) {
      return mid1
    }
    if (arr[mid2] == target) {
      return mid2
    }

    if (target < arr[mid1]) {
      return ternarySearch(left, mid1 - 1, target, arr)
    } else if (target > arr[mid2]) {
      return ternarySearch(mid2 + 1, right, target, arr)
    } else {
      return ternarySearch(mid1 + 1, mid2 - 1, target, arr)
    }
  }

  return -1
}

let arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
const left = 0
const right = arr.length - 1
const target = 5
const result = ternarySearch(left, right, target, arr)

console.log(result1)
console.log(result)