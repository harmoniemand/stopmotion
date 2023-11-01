import { defineStore } from "pinia"
export const useCameraStore = defineStore('camera', {
    state: (): {
        Stream: MediaStream | null,
        IsLoading: boolean,
        IsCameraOpen: boolean,
        HasError: boolean,
        LastError: string,
        CameraVideoElement: HTMLVideoElement | null,
    } => ({
        Stream: null,
        IsLoading: true,
        IsCameraOpen: false,
        HasError: false,
        LastError: "",
        CameraVideoElement: null,
    }),
    actions: {
        StartCameraStream(camera: HTMLVideoElement) {
            this.CameraVideoElement = camera;

            const constraints = {
                audio: false,
                video: {
                    facingMode: 'environment'
                }
            }

            navigator.mediaDevices
                .getUserMedia(constraints)
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
        },

        TakePhoto() :string {
            console.log("taking photo")
            
            const canvas: any = document.createElement('canvas');
            canvas.width = 800;
            canvas.height = 600;
            canvas
                .getContext('2d')
                .drawImage(this.CameraVideoElement, 0, 0, canvas.width, canvas.height);

            const data = canvas.toDataURL('image/png');
            return data;
        }
    },
})
