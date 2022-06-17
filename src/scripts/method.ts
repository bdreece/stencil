import type Param from './param';
import type { ParamData } from './param';
import type Type from "./class";

class Method {
    name: string
    about?: string
    longAbout?: string
    params: Param[]
    returns: Type
    constructor(data: MethodData) {

    }
}

export interface MethodData {
    name: string,
    about?: string,
    longAbout?: string,
    params: ParamData[],
    returns: string,
}

export default Method;