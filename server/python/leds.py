import board
import neopixel
import sys


def main(argv):
    action = argv[0]
    nb_leds = int(argv[1])
    brightness = float(argv[2])

    pixels = neopixel.NeoPixel(board.D18, nb_leds, brightness=brightness, auto_write=False)

    pixels.fill(0)

    if action == "display":
        n = len(sys.argv) - 1
        idx = 0

        for i in range(3, n):
            colors = argv[i].split(",")

            pixels[idx] = (int(colors[0]), int(colors[1]), int(colors[2]))
            idx += 1

    pixels.show()

    return 0


if __name__ == "__main__":
    main(sys.argv[1:])
