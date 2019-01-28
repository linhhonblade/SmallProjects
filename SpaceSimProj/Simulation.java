import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.io.File;
import java.util.Scanner;
import java.lang.Exception;


public class Simulation {
    public ArrayList<Item> loadItems(String fileName) throws Exception{
        System.out.println("[INFO] Load items");
        ArrayList<Item> items = new ArrayList();
        File file = new File(fileName);
        Scanner scan = new Scanner(file);
        while(scan.hasNext()){
            String[] str = scan.nextLine().split("=");
            Item item = new Item(str[0], Integer.parseInt(str[1]));
            items.add(item);
        }
        System.out.println("[INFO] Load items DONE");
        return items;
    }

    public ArrayList<Rocket> loadU1(ArrayList<Item> items){
        System.out.println("[INFO] Load U1");
        ArrayList<Rocket> listU1 = new ArrayList();
        U1 u1 = new U1();
        int count = items.size();
        while(count!=0){
            while(u1.canCarry(items.get(count-1))){
                u1.carry(items.get(count-1));
                count--;
                if(count==0)
                    break;
            }
            listU1.add(u1);
            u1 = new U1();
        }
        if(u1.cargoWeight != 0){
            listU1.add(u1);
        }
        System.out.println("[INFO] Load U1 DONE");
        if(listU1.isEmpty()){
            System.out.println("Cannot construct a suitable fleet of U1s with specified phase");
            return null;
        }
        return listU1;
    }

    public ArrayList<Rocket> loadU2(ArrayList<Item> items){
        System.out.println("[INFO] Load U2");
        ArrayList<Rocket> listU2 = new ArrayList();
        U2 u2 = new U2();
        int count = items.size();
        while(count!=0){
            while(u2.canCarry(items.get(count-1))){
                u2.carry(items.get(count-1));
                count--;
                if(count == 0)
                    break;
            }
            listU2.add(u2); 
            u2 = new U2();
        }
        if(u2.cargoWeight != 0){
            listU2.add(u2);
        }
        System.out.println("[INFO] Load U2 DONE");
        if(listU2.isEmpty()){
            System.out.println("Cannot construct a suitable fleet of U2s with specified phase");
            return null;
        }
        return listU2;
    }

    public int runSimulation(ArrayList<Rocket> rockets){
        System.out.println("[INFO] Simulation start...");
        int count = rockets.get(0).cost;
        for(int i = 0; i < rockets.size(); i++){
            while(!(rockets.get(i).launch() && rockets.get(i).land())){
                count+=count;
                //System.out.println(count);
            }
            count+=count;
        }
        System.out.println("[INFO] Simulation end.");
        return count;
    }
}
