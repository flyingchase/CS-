
public class SelectSort {
    /*
     * 思想: 1.找出数组中最小元素并与第一个元素交换; 2.在剩下中找到最小元素并与数组第二个元素交换,直至整个数组排序;
     * 
     */

    public static void selectionSort(int[] array) {
        int len = array.length;
        for (int i = 0; i < len; i++) {
            for (int item : array)
                System.out.print(item + " ");
            System.out.println();
            int min_index = i;
            for (int j = i + 1; j < len; j++) {
                if (array[j] < array[min_index]) {
                    min_index = j;
                }
            }
            int temp = array[min_index];
            array[min_index] = array[i];
            array[i] = temp;
        }
    }

    public static void main(String[] args) {
        int[] unsortedArray = new int[] { 1, 0, 9, 2, 8, 3, 4, 7, 6, 5, 12, 11 };
        selectionSort(unsortedArray);
        for (int item : unsortedArray)
            System.out.print(item + " ");
        System.out.println();
    }
}
