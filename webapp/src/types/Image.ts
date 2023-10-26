import { Guid } from "guid-typescript";

class Image {
    BASE64: string = "";
    Id: string = "";

    constructor(b64: string, id: string  = '' ) {
        this.Id = id === "" ? Guid.create().toString() : id;
        this.BASE64 = b64;        
    }

}

export default Image;