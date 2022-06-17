import type Type from './type';
import type Method from "./method";

interface Interface {
    name: string,
    about?: string,
    longAbout?: string,
    associated?: Type[],
    implements?: Interface[],
    methods: Method[],
}

export default Interface;