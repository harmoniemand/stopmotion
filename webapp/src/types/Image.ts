import { Guid } from "guid-typescript";

class Image {
    CreatedAt: string = "";
    Id: string = "";
    Url: string = "";

    constructor(img: any) {
        this.Id = img.id === "" ? Guid.create().toString() : img.id;
        this.CreatedAt = img.created_at;
        this.Url = img.url;
    }

}

export default Image;