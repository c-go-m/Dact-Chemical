package constant

//types Errors
const TypeGeneralError = "*useError.GeneralError"
const TypeBussinesError = "*useError.BussinesError"

//Generals Errors
const ErrConnectionDataBase = "ERROR: Problema al iniciar o conectar con la base de datos"
const ErrConnectionStorage = "ERROR: Problema al iniciar o conectar con el Storage"
const ErrEnvironmentVariable = "ERROR: Problema al obtener las variables de entorno"
const ErrInitServer = "ERROR: Problema al Iniciar el servidor"
const ErrDataInvalid = "ERROR: La peticion no es valida"
const ErrSaveFile = "ERROR: No es posible guardar el archivo"

//Bussines Errors
const CodKeyDuplicate = "23505"
const ErrDuplicateRegistry = "ERROR: Registro Duplicado"
const ErrCreateRegistry = "ERROR: No ha sido posible crear el registro en la base de datos"
const ErrDeleteRegistry = "ERROR: No ha sido posible eliminar el registro en la base de datos"
const ErrFindRegistry = "ERROR: No ha sido posible consultar la base de datos"
const ErrUpdateRegistry = "ERROR: No ha sido posible actualizar el registro en la base de datos"

//Codigo Error
const ErrBadRequest = 400
const ErrInternalServerError = 500

//File System
const PathImgSystem = "assets/img/"
const PathBanner = "banner/"
const PathCategory = "category/"
const PathProduct = "product/"
const PathPresentation = "presentation/"
const PathInformation = "information/"
