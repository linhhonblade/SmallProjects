# Notes

- Handler -> Business [-> Repository] -> Storage (tầng trên dùng interface của tầng dưới)
- 2 layer cannot perform unit test: Handler, Storage (only can integration test)
- Business (implement logics) can unit test
- Interface: tui không biết bạn là ai, nhưng tôi biết bạn có thể làm được cái tôi cần
- 