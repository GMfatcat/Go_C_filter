import cv2
import numpy as np
import os

def resize_img(original_image: np.array, scale: float or int) -> np.array:
	new_width = int(original_image.shape[1] * scale)
	new_height = int(original_image.shape[0] * scale)
	resized_image = cv2.resize(original_image, (new_width, new_height))
	return resized_image

def main(filename: str, output_filename: str):
	raw_img = cv2.imread(filename)
	new_img = resize_img(raw_img,5)
	cv2.imwrite(output_filename, new_img)

if __name__ == '__main__':
	img_file = "cat.jpg"
	output_file = os.path.join("large_jpg/",f"resize_{img_file}")
	main(filename = img_file,output_filename = output_file)