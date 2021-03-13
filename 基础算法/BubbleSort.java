public class BubbleSort {
    //冒泡排序 平均/最坏情况均为O(n^2) 空间复杂度O(1)
    public static void bubbleSort(int[] array) {
        int len=array.length;
        for(int i=0;i<len;i++){
            for (int item : array) System.out.print(item+" ");
            System.out.println();
            for(int j=1;j<len-i;j++){
                if(array[j-1]>array[j]){
                    int temp=array[j-1];
                    array[j-1]=array[j];
                    array[j]=temp;
                }
            }
        }
    }
    public static void main(String[] args) {
        int[] unsortedArray = new int[]{6,5,4,3,1,2,8,7,9,0};
        bubbleSort(unsortedArray);
        for (int item : unsortedArray) System.out.print(item+" ");
        System.out.println();
    }
}
