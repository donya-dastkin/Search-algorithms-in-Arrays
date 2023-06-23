//left => low
//right => high
const arr = [1, 2, 29, 33, 55, 86, 95, 120, 220, 301, 450]
let target = 220

function BinarySearch(arr, target) {
  let left = 0
  let right = arr.length - 1
  while (left <= right) {
    let mid = Math.floor((left + right) / 2)
    if (arr[mid] === target)
      return mid
    else if (arr[mid] < target)
      left = mid + 1
    else
      right = mid - 1
  }
  return -1
}

let result = BinarySearch(arr,target)
console.log("result :", result)