public class MergeSort {
    /*
     * 归并排序: 通常用于将两个有序对的数组归并成一个更大的有序数组; 典型的分治应用——>递归排序; 时间复杂度O(NlogN)
     * 空间复杂度O(n)(使用了等长数组);
     */

    // 原地归并如下

    public static void merge(int[] array, int low, int mid, int high) {
        int[] helper = new int[array.length];
        // copy array to helper数组
        for (int k = low; k <= high; k++)
            helper[k] = array[k];

        // 对array[low...mid] 和 array[mid+1...high]进行归并
        int i = low, j = mid + 1;
        for (int k = low; k <= high; k++) {
            // k为当前位置
            if (i > mid) {
                // 左侧部分没有元素
                array[k] = helper[j];
                j++;
            } else if (j > high) {
                // 右侧无元素
                array[k] = helper[i];
                i++;
            } else if (helper[i] > helper[j]) {
                // 右侧最小
                array[k] = helper[j];
                j++;
            } else {
                // 左侧最小
                array[k] = helper[i];
                i++;
            }
        }

    }

    public static void sort(int[] array, int low, int high) {
        if (high <= low)
            return;
        int mid = low + (high - low) / 2;
        sort(array, low, mid);
        sort(array, mid + 1, high);
        merge(array, low, mid, high);
        for (int item : array)
            System.out.print(item + " ");
        System.out.println();
    }

    public static void mergeSort(int[] array) {
        sort(array, 0, array.length - 1);
    }

    public static void main(String[] args) {
        int[] unsortedArray = new int[] { 5, 7, 3, 8, 2, 9, 1, 0, 11, 6, 12 };
        mergeSort(unsortedArray);
        System.out.println("排序后结果是: ");
        for (int item : unsortedArray)
            System.out.print(item + " ");
    }

}
