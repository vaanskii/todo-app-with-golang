import time

class Owner:
    def __init__(self, name):
        self.name = name

class GasEngine:
    def __init__(self, mpg, gallons, owner_info):
        self.mpg = mpg
        self.gallons = gallons
        self.owner_info = owner_info

def main():
    start_time = time.time()  # Start the timer

    # Create an instance of GasEngine
    my_engine = GasEngine(25, 15, Owner("Giorgi"))

    # Print the attributes
    print(my_engine.mpg, my_engine.gallons, my_engine.owner_info.name)

    elapsed_time = (time.time() - start_time) * 1000 
    print(f"Execution time: {elapsed_time:.6f} ms")  

if __name__ == "__main__":
    main()
