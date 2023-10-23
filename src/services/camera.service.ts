import Image from "../types/Image";
import CameraEvent from "../Events/CameraEvent";


class CameraService extends EventTarget {

    private static instance: CameraService;

    Stream: MediaStream | null = null;
    LastImage: Image | null = null;

    IsLoading: boolean = true;
    IsCameraOpen: boolean = false;

    HasError: boolean = false;
    LastError: string = "";

    CameraVideoElement: HTMLVideoElement | null = null;

    // private _onNewImage: CameraEvent = new CameraEvent("onNewImage");
    private _onCameraError: Event = new Event("onCameraError");

    constraints = {
        audio: false,
        video: {
            facingMode: 'environment'
        }
    }

    private constructor() {
        super();
    }

    static getInstance() {
        if (!CameraService.instance) {
            CameraService.instance = new CameraService();
        }
        
        return CameraService.instance;
    }

    StartCameraStream(camera: HTMLVideoElement) {
        this.CameraVideoElement = camera;

        navigator.mediaDevices
            .getUserMedia(this.constraints)
            .then(stream => {
                this.HasError = false;

                if (this.CameraVideoElement) {
                    console.log("Stream started")
                    this.IsLoading = false;
                    this.IsCameraOpen = true;

                    this.CameraVideoElement.srcObject = stream;
                } else {
                    console.log("cameta not found")
                }
            })
            .catch(error => {
                this.IsLoading = false;
                console.log(error);
            });
    }


    TakePhoto() {
        const canvas: any = document.createElement('canvas');
        canvas.width = 800;
        canvas.height = 600;
        canvas
            .getContext('2d')
            .drawImage(this.CameraVideoElement, 0, 0, canvas.width, canvas.height);

        const data = canvas.toDataURL('image/jpg');
        this.LastImage = new Image(data);

        console.log("image taken and saved to last images");

        this.dispatchEvent(new CameraEvent("onNewImage", this.LastImage));
    }
}

export default CameraService;