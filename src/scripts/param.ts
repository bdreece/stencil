import type Type from "./type";

class Param {
    constructor(data: ParamData) {

    }

    public get value(): ParamData {

    }
}

export interface ParamData {
    name: string,
    about?: string,
    longAbout?: string,
    type?: string;
}

export default Param;