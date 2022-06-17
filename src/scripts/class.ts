import type Method from './method';
import type Param from './param';

class Class {
    name: string
    about?: string
    longAbout?: string
    methods: Method[]
    members: Param[]
    constructor(data) {

    }

    public get value(): ClassData {

    }
}

export interface ClassData {
    name: string,
    about?: string,
    longAbout?: string,
    methods: Method[],
    members: Param[],
}

export default Class;