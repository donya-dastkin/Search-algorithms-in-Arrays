const linearSearch = (arr, target) => {
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === target) return i;
  }
  return -1;
};

const arr =[1,2,3,5,7,6,8,10,9]
const target = 6
const Result = linearSearch(arr,target)
console.log("Result :" , Result)