import java.util.ArrayList;


public class Main {
    public static void main(String[] args) throws Exception{
        Simulation mySimulation = new Simulation();
        ArrayList<Item> phase1 = mySimulation.loadItems("phase-1.txt");
        ArrayList<Item> phase2 = mySimulation.loadItems("phase-2.txt");
        ArrayList<Rocket> u1Phase1 = mySimulation.loadU1(phase1);
        ArrayList<Rocket> u1Phase2 = mySimulation.loadU1(phase2);
        ArrayList<Rocket> u2Phase1 = mySimulation.loadU2(phase1);
        ArrayList<Rocket> u2Phase2 = mySimulation.loadU2(phase2);

        if(u1Phase1!=null)
        	System.out.println("U1 in phase 1: " + mySimulation.runSimulation(u1Phase1) + " Millions");
        if(u1Phase2!=null)
        	System.out.println("U1 in phase 2: " + mySimulation.runSimulation(u1Phase2) + " Millions");
        if(u2Phase1!=null)
	        System.out.println("U2 in phase 1: " + mySimulation.runSimulation(u2Phase1) + " Millions");
	    if(u2Phase2!=null)
        	System.out.println("U2 in phase 2: " + mySimulation.runSimulation(u2Phase2) + " Millions");

    }
}
