import Image from "../types/Image";

class CameraEvent extends Event {
    Image: Image | null = null;

    constructor(type: string, image: Image | null) {
        super(type);
        this.Image = image;
    }
}

export default CameraEvent;