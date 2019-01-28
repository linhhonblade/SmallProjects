public class U2 extends Rocket {
    public U2(){
        super(100,10000,18000);
    }
    public boolean launch(){
        double chanceOfExplosion = 4*(this.cargoWeight/(this.maxWeight-this.weight));
        double getChance = Math.random()*100 - chanceOfExplosion;
        if(getChance <= 0)
            return false;
        else
            return true;
    }

    public boolean land(){
        double chanceOfCrash = 8*(this.cargoWeight/(this.maxWeight-this.weight));
        double getChance = Math.random()*100 - chanceOfCrash;
        if(getChance <= 0)
            return false;
        else
            return true;
    }
}
