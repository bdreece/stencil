import type Type from './type';

class Component {
    name: string
    about?: string
    longAbout?: string
    types: Type[]

    constructor(data: ComponentData) {
        this.name = data.name;
        this.about = data.about;
        this.longAbout = data.longAbout;
        this.types = 
    }

    public get value(): ComponentData {
        return {
            name: this.name,
            about: this.about,
            longAbout: this.longAbout,
            types: this.types,
        };
    }
}

export interface ComponentData {
    name: string,
    about?: string,
    longAbout?: string,
    types: Type[]
}

export default Component;