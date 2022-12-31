class Sensor:
    x: int = 0
    y: int = 0

    def update(self) -> None:
        self.x += 1
        self.y += 1

        if self.x >= 100000:
            self.x = 0
        if self.y >= 100000:
            self.y = 0

    def get_coords(self) -> tuple[float, float]:
        return (self.x, self.y)
