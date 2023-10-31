import Project from "./Project";

class AppState {
    ProjectId: string = "";

    constructor(d: any) {
        this.ProjectId = d?.ProjectId ? d.ProjectId : "";
    }
}

export default AppState;