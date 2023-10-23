class AppState {
    SessionId: string = "";
    Name: string = "";
    Description: string = "";
    Authors: string = "";

    constructor(d: any) {
        this.SessionId = d?.sessionId ? d.sessionId : "";
        this.Name = d?.name ? d.name : "";
        this.Description = d?.description ? d.description : "";
        this.Authors = d?.authors ? d.authors : "";
    }
}

export default AppState;