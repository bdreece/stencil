import toml from 'toml';

import Component from "./component";
import type { ComponentData } from "./component";

class Project {
    name: string
    about?: string
    longAbout?: string
    language: string
    catalog: Component[]
    components: Component[]

    constructor(data) {
        this.name = data.name;
        this.about = data.about;
        this.longAbout = data.longAbout;
        this.language = data.language;
        if (data.catalog) {
            data.catalog.forEach(component => {
                this.catalog.push(new Component(component))
            })
        }
        if (data.components) {
            data.components.forEach(component => {
                this.components.push(new Component(component))
            })
        }
    }

    public async fetchPattern(url: string) {
        const response = await fetch(url)
        const text = await response.text()
        const component = new Component(toml.parse(text))
        this.catalog.push(component)
    }

    public get value(): ProjectData {
        return {
            name: this.name,
            about: this.about,
            longAbout: this.longAbout,
            language: this.language,
            catalog: this.catalog.map(component => component.value),
            components: this.components.map(component => component.value),
        }
    }
}

export interface ProjectData {
    name: string,
    about?: string,
    longAbout?: string,
    language: string,
    catalog?: ComponentData[],
    components?: ComponentData[],
}

export default Project;