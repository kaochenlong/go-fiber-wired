import { application } from "./application"

// TODO: load all controller files automatically
import HelloController from "./hello_controller"
application.register("hello", HelloController)
