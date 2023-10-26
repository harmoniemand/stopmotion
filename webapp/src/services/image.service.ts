
import CameraEvent from "../Events/CameraEvent";
import Image from "../types/Image";
import CameraService from "./camera.service";
import { ref } from "vue";

class ImageService {

    private static instance: ImageService;

    Images = ref<Image[]>([]);

    private cameraService: CameraService | null = null;;

    
    private constructor() {
        this.cameraService = CameraService.getInstance();
        this.cameraService.addEventListener('onNewImage', (e: any) => {
            const img = (e as CameraEvent).Image;

            if (img) {
                console.log("new image", img);
                this.Add(img);
            }
        })
    }

    static getInstance() {
        if (!ImageService.instance) {
            ImageService.instance = new ImageService();
        }
        
        return ImageService.instance;
    }

    Add(image: Image) {
        console.log("add image", image)
        this.Images.value.push(image);
    }

    Delete(image: Image) {
        this.Images.value = this.Images.value.filter(i => i.Id !== image.Id);
    }

    Reset() {
        this.Images.value = [];
    }
}

export default ImageService;