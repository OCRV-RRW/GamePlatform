import { Preview } from "./preview_type"

export type Game = {
    description: string,
    title: string,
    icon: string,
    previews: Array<Preview>,
    src: string,
    id: string
}