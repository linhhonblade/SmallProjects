public class U1 extends Rocket {
    public U1(){
        super(120,18000,29000);
    }
    public boolean launch(){
        double chanceOfExplosion = 5*(this.cargoWeight/(this.maxWeight-this.weight));
        double getChance = Math.random()*100 - chanceOfExplosion;
        if(getChance <= 0)
            return false;
        else
            return true;
    }

    public boolean land(){
        double chanceOfCrash = 1*(this.cargoWeight/(this.maxWeight-this.weight));
        double getChance = Math.random()*100 - chanceOfCrash;
        if(getChance <= 0)
            return false;
        else
            return true;
    }
}
