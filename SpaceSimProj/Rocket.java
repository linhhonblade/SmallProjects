public abstract class Rocket implements SpaceShip {
    protected int cost;
    protected int weight;
    protected int maxWeight;
    protected int cargoWeight;

    Rocket(int cost, int weight, int maxWeight){
        this.cost = cost;
        this.weight = weight;
        this.maxWeight = maxWeight;
        this.cargoWeight = 0;
    }

    public boolean launch(){
        return true;
    }
    public boolean land(){
        return true;
    }


    public final boolean canCarry(Item item){
        if(this.weight + this.cargoWeight + item.getWeight() <= this.maxWeight)
            return true;
        else
            return false;
    }


    public final void carry(Item item){
        this.cargoWeight+=item.getWeight();
    }
}
