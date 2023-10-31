import Image from "./Image";

class Project {
    id: string = "";
    name: string = "";
    description: string = "";
    authors: string = "";
    images: Image[] = [];

    constructor(d: any) {
        this.id = d?.id ? d.id : "";
        this.name = d?.name ? d.name : "";
        this.description = d?.description ? d.description : "";
        this.authors = d?.authors ? d.authors : "";

        this.images = [];
        if (d?.images) {
            d.images.forEach((i: any) => {
                this.images.push(new Image(i))
            })
        }
    }
}

export default Project;