import type Class from './class';
import type Interface from './interface';

type PrimitiveType = "void" | "integer" | "float" | "char" | "string"

interface Custom {
    name: string
    inner?: Type
}

type ObjectType = Class | Interface | Custom;

type Type = PrimitiveType | ObjectType;

export default Type;