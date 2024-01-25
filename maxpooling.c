// maxpooling.c

#include <stdint.h>

// Max pooling 函式
void max_pooling(const uint8_t* input, uint8_t* output, int width, int height, int pool_size) {
    int new_width = width / pool_size;
    int new_height = height / pool_size;

    for (int y = 0; y < new_height; y++) {
        for (int x = 0; x < new_width; x++) {
            uint8_t max_value = 0;
            for (int i = 0; i < pool_size; i++) {
                for (int j = 0; j < pool_size; j++) {
                    uint8_t value = input[(y * pool_size + i) * width + (x * pool_size + j)];
                    if (value > max_value) {
                        max_value = value;
                    }
                }
            }
            output[y * new_width + x] = max_value;
        }
    }
}
