/* Selection Sort - select the minimum and swap with i
13 4 7 9 -> (1) 4 13 7 9 -> (2) 4 7 13 9 -> (3) 4 7 9 13 : 4 elements sorted in 3 steps 

Bubble Sort - compare adjacent elements and swap

Insertion Sort - insert the element in the correct position in the sorted array
*/
function selectionSort(arr, n) {
    for (let i = 0; i <= n - 2; i++) {
      let min = i;
      for (let j = i; j <= n - 1; j++) {
        if (arr[j] < arr[min]) {
          min = j;
        }
      }
      swap(arr[i], arr[min]);
    }
  }
  
  function bubbleSort(arr, n) {
    for (let i = n - 1; i >= 1; i++) {
      let didSwap = 0;
      for (let j = 0; j <= i - 1; j++) {
        if (arr[j] > arr[j + 1]) {
          swap(arr[j], arr[j + 1]);
          didSwap = 1;
        }
      }
      if (didSwap == 1) {
        break;
      }
    }
  }
  function insertionSort(arr, n) {
    for (let i = 0; i <= n - 1; i++) {
      let j = i;
      while (j > 0 && arr[j - 1] > arr[j]) {
        swap(arr[j - 1], arr[j]);
        j--;
      }
    }
  }