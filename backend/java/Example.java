class Example {
  public static void main(String[] args) throws java.io.IOException {
    int len = 10;
    int[] list = new int[len];
    for (int i = 0; i < len; i++) {
      list[i] = i * i;
    }
    for (int v : list) {
      System.out.print(v + " ");
    }
  }
}

class TwoD {
}