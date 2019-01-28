public interface SpaceShip {
    boolean launch();
    boolean land();
    abstract boolean canCarry(Item item);
    abstract void carry(Item item);
}
