public class InsertionSort {
    /*
     * 插入排序: 对未排序序列从后往前扫描 找到相应位置插入;
     * 
     */

    public static void insertionSort(int[] array) {
        int len = array.length;

        for (int i = 0; i < len; i++) {
            int index = i, array_i = array[i];
            while (index > 0 && array[index - 1] > array_i) {
                array[index] = array[index - 1];
                index -= 1;
            }
            array[index] = array_i;
            System.out.printf("这是第%d次排序后结果: ", i + 1);
            for (int item : array)
                System.out.print(item + " ");
            System.out.println();
        }
    }

    public static void main(String[] args) {
        int[] unsortedArray = new int[] { 9, 5, 1, 2, 8, 3, 7, 10, 6, 0, 11, 2 };
        insertionSort(unsortedArray);
        for (int item : unsortedArray)
            System.err.print(item + " ");
        System.out.println();

    }

}
