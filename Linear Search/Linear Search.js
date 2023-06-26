const linearSearch = (arr, x) => {
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === x) return i;
  }
  return -1;
};

const arr =[1,2,3,5,7,6,8,10,9]
const x = 6
const Result = linearSearch(arr,x)
console.log("Result :" , Result)