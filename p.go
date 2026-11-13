
package week6;
import java.util.LinkedList;
import java.util.ArrayList;
import week6.Car;
/**
 *
 * @author W0319
 */
public class Main {
    
 
    public static void main(String[] args) {
        
        int[] arrayInt = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        for (int i = 0; i < arrayInt.length; i++) {
            System.out.println("data ke - "+ arrayInt[i]);
        }
        System.out.println("");
        
        
        double[] arrayDouble = new double[5];
        arrayDouble[0] = 28.5;
        arrayDouble[1] = 25.5;
        arrayDouble[2] = 26.5;
        arrayDouble[3] = 27.5;
        System.out.println("data ke - "+ arrayDouble[0]);
        System.out.println("data ke - "+ arrayDouble[1]);
        System.out.println("data ke - "+ arrayDouble[2]);
        System.out.println("data ke - "+ arrayDouble[3]);
        System.out.println("");
        for (int i = 0; i < arrayDouble.length; i++) {
            System.out.println("data ke - "+ arrayDouble[i]);
        }
        System.out.println("");
        
        String[] arrayString = new String[5];
        arrayString[0] = "Nama 0";
        arrayString[2] = "Nama 2";
        for (int i = 0; i < arrayString.length; i++) {
            System.out.println("data ke - "+ arrayString[i]);
        }
        System.out.println("");
        
        
        ArrayList<String> fruits = new ArrayList<>();
        //add
        fruits.add("Apple");
        fruits.add("Banana");
        fruits.add("Cherry");
        //remove
        fruits.remove(1);
        //get
        System.out.println("First fruit: " + fruits.get(0));
        //loop
        for (String fruit : fruits) {
            System.out.println(fruit);
        }
        System.out.println("");
        
        System.out.println("array List of object");
        ArrayList<Car> arrayListCar = new ArrayList<>();
        arrayListCar.add(new Car("Honda", "BR-V", 2025));
        System.out.println(arrayListCar.get(0).brand + " " + arrayListCar.get(0).type + " ");
        System.out.println("");
        
        LinkedList<Integer> numbers = new LinkedList<>();
        //add
        numbers.add(10);
        numbers.add(20);
        numbers.add(30);
        //remove
        numbers.remove(1);
        //get
        System.out.println("First element: " + numbers.getFirst());
        System.out.println("Last element: " + numbers.getLast());
        //loop
        for (int number : numbers) {
            System.out.println(number);
        }
        System.out.println("");
        
        System.out.println("Linked list of object");
        LinkedList<Car> LinkedListCar = new LinkedList<>();
        LinkedListCar.add(new Car("Honda", "C-RV", 2025));
        System.out.println(LinkedListCar.get(0).brand + " " + LinkedListCar.get(0).type + " ");
        System.out.println("");
        
        
        Car[] arrayCar = {
            new Car("Toyota", "Fortuner", 2020),
            new Car("Honda", "BR-V", 2022),
            new Car("Suzuki", "Jimny", 2020),
        };
        for (Car car : arrayCar) { 
            System.out.println(car.brand + " " + car.type + " " + String.valueOf(car.year));
        }
        
    }   
    }
    