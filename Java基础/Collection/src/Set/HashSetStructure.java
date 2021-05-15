package Set;

@SuppressWarnings({"all"})
public class HashSetStructure {
    public static void main(String[] args) {
        // 数组+链表模拟HashSet/HashMap structure
        // 数组类型为 Node
        Node[] tables = new Node[16];
        System.out.println(tables);

        Node wo = new Node("ai", null);
        tables[2]=wo;
        Node wenlei = new Node("wlei", null);
//        tables[3]=wenlei;
        wo.next=wenlei;

        Node zhou = new Node("song", null);
        wenlei.next=zhou;

        Node wang = new Node("xu", null);
        tables[3]=wang;

        System.out.println("tables = " + tables);


    }
}
class Node{
    // Node point to Node--->linklist
    Object item;
    Node next;

    public Node(Object item, Node next) {
        this.item = item;
        this.next = next;
    }
}
