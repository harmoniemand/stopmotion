import { Guid } from "guid-typescript";

class Image {
    data: string = "";
    created_at: string = "";
    Id: string = "";

    constructor(img: any) {
        this.Id = img.id === "" ? Guid.create().toString() : img.id;
        this.data = img.data;
        this.created_at = img.created_at;
    }

}

export default Image;