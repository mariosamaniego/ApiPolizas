--Script Tablas
--Tablas
----------------------------------------------------------------------------------------
--=========================================================================
--Elaboro  : Mario Samaniego
--Descripcion: se crea tabla de inventario
---------------------------------------------------------------------------
DO
$$
DECLARE
iExiste integer;
BEGIN


SELECT count(*) INTO iExiste FROM pg_tables WHERE tablename= 'ctl_inventario';

IF (iExiste = 0) THEN

	CREATE TABLE IF NOT EXISTS public.ctl_inventario
	(
		sku integer NOT NULL DEFAULT 0,
		nombre CHARACTER(30) COLLATE pg_catalog."default" NOT NULL DEFAULT ' '::bpchar,
		cantidad integer NOT NULL DEFAULT 0,
		CONSTRAINT pk_ctl_inventario PRIMARY KEY (sku)
	);
	CREATE INDEX idx_ctl_inventario ON ctl_inventario (sku);
	
	GRANT ALL ON TABLE ctl_inventario TO postgres;

	COMMENT ON TABLE ctl_inventario IS 'Se crea catalogo de inventario';

	COMMENT ON COLUMN ctl_inventario.sku IS 'Guarda el número del sku';

	COMMENT ON COLUMN ctl_inventario.nombre IS 'Guarda el nombre del sku';
	
	COMMENT ON COLUMN ctl_inventario.cantidad IS 'Guarda la cantidad de unidades del articulo';
		  
END IF;

END;

$$


---------------------------------------------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------------------------------------
 --=========================================================================
-- Elaboro  : Mario Samaniego
--Descripcion: se crea tabla de Empleado
---------------------------------------------------------------------------
DO
$$
DECLARE
iExiste integer;
BEGIN


SELECT count(*) INTO iExiste FROM pg_tables WHERE tablename= 'ctl_empleado';

IF (iExiste = 0) THEN

	CREATE TABLE IF NOT EXISTS public.ctl_empleado
	(
		id_empleado integer NOT NULL DEFAULT 0,
		nombre CHARACTER(30) COLLATE pg_catalog."default" NOT NULL DEFAULT ' '::bpchar,
		apellido CHARACTER(30) COLLATE pg_catalog."default" NOT NULL DEFAULT ' '::bpchar,
		puesto CHARACTER(30) COLLATE pg_catalog."default" NOT NULL DEFAULT ' '::bpchar,
		CONSTRAINT pk_ctl_empleado PRIMARY KEY (id_empleado)
	);
	CREATE INDEX idx_ctl_empleado ON ctl_empleado (id_empleado);
	
	GRANT ALL ON TABLE ctl_empleado TO postgres;

	COMMENT ON TABLE ctl_empleado IS 'Se crea catalogo de empleado';

	COMMENT ON COLUMN ctl_empleado.id_empleado IS 'Guarda el número del empleado';

	COMMENT ON COLUMN ctl_empleado.nombre IS 'Guarda el nombre del empleado';
	
	COMMENT ON COLUMN ctl_empleado.apellido IS 'Guarda el apellido del empleado';
	
	COMMENT ON COLUMN ctl_empleado.puesto IS 'Guarda el puesto del empleado';
		  
END IF;

END;

$$

------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
 --=========================================================================
-- Elaboro  : Mario Samaniego
--Descripcion: se crea tabla de Polizas
---------------------------------------------------------------------------
DO
$$
DECLARE
iExiste integer;
BEGIN


SELECT count(*) INTO iExiste FROM pg_tables WHERE tablename= 'ctl_polizas';

IF (iExiste = 0) THEN

	CREATE TABLE IF NOT EXISTS public.ctl_polizas
	(
		id_polizas 	INTEGER NOT NULL DEFAULT 0,
		empleado_genero INTEGER NOT NULL DEFAULT 0,
		sku 		INTEGER NOT NULL DEFAULT 0,
		cantidad 	INTEGER NOT NULL DEFAULT 0,
		fecha 		DATE NOT NULL DEFAULT NOW()::DATE,
		CONSTRAINT pk_ctl_polizas PRIMARY KEY (id_polizas)
	);
	CREATE INDEX idx_ctl_polizas ON ctl_polizas (id_polizas);
	
	GRANT ALL ON TABLE ctl_polizas TO postgres;

	COMMENT ON TABLE ctl_polizas IS 'Se crea catalogo de empleado';

	COMMENT ON COLUMN ctl_polizas.id_polizas IS 'Guarda el número de la poliza';

	COMMENT ON COLUMN ctl_polizas.empleado_genero IS 'Guarda el nombre del empleado que genero la poliza';
	
	COMMENT ON COLUMN ctl_polizas.sku IS 'Guarda el numero del sku del articulo';
	
	COMMENT ON COLUMN ctl_polizas.cantidad IS 'Guarda la cantidad de poliza';
	
	COMMENT ON COLUMN ctl_polizas.fecha IS 'Guarda la fecha de la poliza';
		  
END IF;

END;

$$

------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
--INSERT CTL_EMPLEADO
INSERT INTO ctl_empleado VALUES 
(98493701,'MARIO ALBERTO','SAMANIEGO VEL','PLOMERO'),
(90215021,'LUIS','SALAZAR','AYUDANTE'),
(93825994,'RENE','LLAVE','VELADOR'),
(98463187,'PANCHO','CANICAS','GERENTE'),
(97848591,'ERIC','SON','VENDEDOR')

--INSERT ctl_inventario
INSERT INTO ctl_inventario VALUES
(489559,'ORGANIZADOR BIN',100),
(596728,'BICIBLETA TURBO',100),
(476343,'ACCESORIO SATAR',100),
(362603,'COMEDOR ZAZ',100),
(644838,'MINISPLIT MABE',100),
(506259,'CASCO PAWPATROL',100),
(456375,'CORTINA STARHAUS',100),
(234817,'DIADEMA TURTLE',100),
(645036,'VENTILADOR DREAMETCH',100),
(547557,'LLANTA STARPRIX',100)


--INSERT ctl_polizas
INSERT INTO ctl_polizas VALUES
(23546,98493701,596728,1,'2023-07-03'),
(23547,90215021,362603,1,'2023-07-05'),
(23548,93825994,506259,1,'2023-07-04'),
(23549,98463187,645036,1,'2023-07-03'),
(23550,97848591,547557,1,'2023-07-02'),
(23551,98493701,644838,1,'2023-07-01')


------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
SELECT idempleado,nombre,apellido,puesto FROM fun_consultaempleados()
CREATE OR REPLACE FUNCTION fun_consultaempleados()
  RETURNS TABLE(idempleado INTEGER,nombre CHARACTER, apellido CHARACTER, puesto CHARACTER)AS
$BODY$
-----------------------------------------------------------------------------------------------
--Funcion para informacion de los empleados
-----------------------------------------------------------------------------------------------
DECLARE  
	
BEGIN

	RETURN QUERY
	SELECT a.id_empleado,a.nombre,a.apellido ,a.puesto
	FROM ctl_empleado as a;
	
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;

----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------
select * from fun_consultaarticulos()

CREATE OR REPLACE FUNCTION fun_consultaarticulos()
  RETURNS TABLE(sku INTEGER,nombre CHARACTER, cantidad INTEGER)AS
$BODY$
-----------------------------------------------------------------------------------------------
--Funcion para consultar el inventario
-----------------------------------------------------------------------------------------------
DECLARE  
	
BEGIN

	RETURN QUERY
	SELECT a.sku,a.nombre,a.cantidad
	FROM ctl_inventario as a;
	
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;


----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------

SELECT id_polizas,empleado_genero, sku, cantidad, nombre,apellido, puesto FROM fun_consultapolizasempleado(98493701)
drop function fun_consultapolizasempleado(INTEGER)

CREATE OR REPLACE FUNCTION fun_consultapolizasempleado(INTEGER)
  RETURNS TABLE(id_polizas INTEGER,empleado_genero INTEGER, sku INTEGER, cantidad INTEGER, nombre CHARACTER,apellido CHARACTER, puesto CHARACTER)AS
$BODY$
-----------------------------------------------------------------------------------------------
--Funcion para mostrar las polizas del empleado
-----------------------------------------------------------------------------------------------
DECLARE  
	iEmpleado ALIAS FOR $1;
	
BEGIN
	
	RETURN QUERY
	SELECT a.id_polizas,a.empleado_genero,a.sku,a.cantidad,b.nombre,b.apellido, b.puesto
	FROM ctl_polizas AS a 
	INNER JOIN ctl_empleado AS b ON b.id_empleado = a.empleado_genero 
	WHERE a.empleado_genero = iEmpleado;
	
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;

----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------
select fun_registraempleado(98493702,'luis','carnalillo','basquetbolista')
drop function fun_registraempleado(INTEGER, CHARACTER VARYING, CHARACTER VARYING, CHARACTER VARYING) 

CREATE OR REPLACE FUNCTION fun_registraempleado(INTEGER, CHARACTER VARYING, CHARACTER VARYING, CHARACTER VARYING) 
RETURNS INTEGER 
AS
$BODY$
-----------------------------------------------------------------------------------------------
--Funcion para agregar un nuevo empleado
-----------------------------------------------------------------------------------------------
DECLARE 
	iIdempleado ALIAS FOR $1;
	cNombre ALIAS FOR $2;
	cApellido ALIAS FOR $3;
	cPuesto ALIAS FOR $4;

BEGIN
	IF iIdempleado IS NULL OR iIdempleado = 0 THEN
		RETURN 0;
	ELSE
		IF NOT EXISTS (SELECT FROM ctl_empleado WHERE id_empleado = iIdempleado) THEN		
			INSERT INTO ctl_empleado 
			(id_empleado, 
			 nombre,
			 apellido,
			 puesto) 
			VALUES 
			(iIdempleado,
			 cNombre,
			 cApellido,
			 cPuesto);
			 RETURN 1;
		END IF;
		RETURN 0;
	END IF;
	
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;

----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------


select id_polizas, empleado_genero, sku, cantidad, nombre, apellido, puesto from fun_consultapoliza(23546)
drop function fun_consultapoliza(INTEGER)
CREATE OR REPLACE FUNCTION fun_consultapoliza(INTEGER)
  RETURNS TABLE(id_polizas INTEGER,empleado_genero INTEGER, sku INTEGER, cantidad INTEGER, nombre CHARACTER, apellido CHARACTER, puesto CHARACTER)AS
$BODY$
-----------------------------------------------------------------------------------------------
--Funcion para informacion de una poliza 
-----------------------------------------------------------------------------------------------
DECLARE  
	iPoliza ALIAS FOR $1;
	
BEGIN

	RETURN QUERY
	SELECT a.id_polizas,a.empleado_genero,a.sku,a.cantidad,b.nombre,b.apellido, b.puesto
	FROM ctl_polizas AS a 
	INNER JOIN ctl_empleado AS b ON b.id_empleado = a.empleado_genero 
	WHERE a.id_polizas = iPoliza;
	
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;

----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------

select fun_registrarpoliza(23554,98493701,123,100)
drop function fun_registrarpoliza(INTEGER, INTEGER, INTEGER, INTEGER) 

CREATE OR REPLACE FUNCTION fun_registrarpoliza(INTEGER, INTEGER, INTEGER, INTEGER) 
RETURNS INTEGER 
AS
$BODY$
-----------------------------------------------------------------------------------------------
--Funcion para agregar una nueva poliza
-----------------------------------------------------------------------------------------------
DECLARE 
	iIdPoliza ALIAS FOR $1;
	iEmpGenero ALIAS FOR $2;
	iSku ALIAS FOR $3;
	iCantidad ALIAS FOR $4;

	iNuevo INTEGER;
	iResultado INTEGER;
	icantidadsku INTEGER;
BEGIN
	IF iIdPoliza IS NULL OR iIdPoliza = 0 THEN
		iNuevo = 0;
	ELSE
		IF iEmpGenero IS NULL OR iEmpGenero = 0 THEN
			iNuevo = 0;
		ELSE
			IF NOT EXISTS (SELECT FROM ctl_empleado WHERE id_empleado = iEmpGenero) THEN
				iNuevo = 2;
			ELSE
				IF NOT EXISTS (SELECT FROM ctl_inventario WHERE sku = iSku) THEN
					
					iNuevo = 3;
				ELSE
					SELECT cantidad INTO icantidadsku FROM ctl_inventario WHERE sku = iSku;
					
					IF iCantidad <=  icantidadsku THEN
						iResultado = icantidadsku - iCantidad;
						IF NOT EXISTS (SELECT FROM ctl_polizas WHERE id_polizas = iIdPoliza) THEN
								INSERT INTO ctl_polizas 
								(id_polizas, 
								 empleado_genero,
								 sku,
								 cantidad) 
								VALUES 
								(iIdPoliza,
								 iEmpGenero,
								 iSku,
								 iCantidad);
								 UPDATE ctl_inventario SET cantidad = iResultado WHERE sku = iSku;
								 iNuevo = 1;
						ELSE
							iNuevo = 5;
						END IF;
					ELSE
						iNuevo = 4;
					END IF;
				END IF;
			END IF;
		END IF;
	END IF;
	RETURN iNuevo;
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;


----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------


select fun_eliminargeneral from fun_eliminargeneral(3,489559)
drop function fun_eliminargeneral(INTEGER, INTEGER ) 

CREATE OR REPLACE FUNCTION fun_eliminargeneral(INTEGER, INTEGER) 
RETURNS INTEGER 
AS
$BODY$
-----------------------------------------------------------------------------------------------
--Eliminar Registro
-----------------------------------------------------------------------------------------------
DECLARE 
	iOpcion ALIAS FOR $1;
	iEliminar ALIAS FOR $2;

	iRespuesta INTEGER;
BEGIN
	IF iOpcion = 1 THEN
		DELETE FROM ctl_polizas WHERE id_polizas = iEliminar;
		iRespuesta = 1;
	ELSIF iOpcion = 2 THEN
		IF NOT EXISTS (SELECT FROM ctl_polizas WHERE empleado_genero = iEliminar) THEN
			DELETE FROM ctl_empleado WHERE id_empleado = iEliminar;
			iRespuesta = 1;
		ELSE
			iRespuesta = 2;--empleado existe en tabla de polizas, no podra eliminar
		END IF;
	ELSIF iOpcion = 3 THEN
		DELETE FROM ctl_inventario WHERE sku =  iEliminar;
		iRespuesta = 1;
	ELSE
		iRespuesta = 0;
	END IF;
	RETURN iRespuesta;
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;


----------------------------------------------------------------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------------------------------------------------------------


Select * from fun_actualizarpoliza(23555,98493701,489559,1,'s','s','s')
Select * from fun_actualizarpoliza(23555,98493701,476343,1,'s','s','s')

489559	"ORGANIZADOR BIN               " pone 3 que corresponde a la poliza
476343	"ACCESORIO SATAR               " quita la cantidad que uses en el parametro
DROP FUNCTION fun_actualizarpoliza(INTEGER,INTEGER,INTEGER,INTEGER,CHARACTER,CHARACTER,CHARACTER)


CREATE OR REPLACE FUNCTION fun_actualizarpoliza(INTEGER,INTEGER,INTEGER,INTEGER,CHARACTER,CHARACTER,CHARACTER)
RETURNS INTEGER 
AS
$BODY$
-----------------------------------------------------------------------------------------------
--Actualizar Poliza
-----------------------------------------------------------------------------------------------
DECLARE 
	iPoliza 	ALIAS FOR $1;
	iEmpleado 	ALIAS FOR $2;
	iSku		ALIAS FOR $3;
	iCantidad	ALIAS FOR $4;
	sNombre		ALIAS FOR $5;
	sApellido	ALIAS FOR $6;
	sPuesto		ALIAS FOR $7;

	iRespuesta	INTEGER;
	iskuActual	INTEGER;
	icantActualPoliza INTEGER;
	iCantActualInv INTEGER;
	icantidadsku INTEGER;
	iResultado INTEGER;
	iResultadoactual INTEGER;
	
	cSql	TEXT;
BEGIN
	SELECT sku INTO iskuActual FROM ctl_polizas WHERE id_polizas = iPoliza;
	SELECT cantidad INTO icantActualPoliza FROM ctl_polizas WHERE id_polizas = iPoliza;
	SELECT cantidad INTO iCantActualInv FROM ctl_inventario WHERE sku = iskuActual;
	
	IF EXISTS (SELECT FROM ctl_inventario WHERE sku = iSku) THEN --si existe el nuevo sku
		SELECT cantidad INTO icantidadsku FROM ctl_inventario WHERE sku = iSku; --calcula el total de articulos  del nuevo sku
			IF iCantidad <=  icantidadsku THEN --si el total es mayor a la cantidad que introdujo
				
				--iResultado es el total - lo que recibe de parametro
				IF iskuActual = iSku THEN
					IF iCantidad <= icantActualPoliza THEN
						iResultado = icantActualPoliza - iCantidad;
						iResultado = icantidadsku + iResultado;
						UPDATE ctl_inventario SET cantidad = iResultado WHERE sku = iSku;
					ELSE
						iResultado = iCantidad - icantActualPoliza;
						iResultado = icantidadsku - iResultado;
						UPDATE ctl_inventario SET cantidad = iResultado WHERE sku = iSku;
					END IF;
					iRespuesta = 1;
				ELSE
					iResultado = icantidadsku - iCantidad; --resultado es igual al total menos la cantidad que introdujo
					iResultadoactual = iCantActualInv + icantActualPoliza; -- al skuactual le tiene que agregar lo que se le habia quitado del inventario
					UPDATE ctl_inventario SET cantidad = iResultadoactual WHERE sku = iskuActual; --Actualizar el inventario viejo sumandole la cantidad del sku (anterior)
					UPDATE ctl_inventario SET cantidad = iResultado WHERE sku = iSku; 			 --Actualizar el inventario del nuevo sku
					iRespuesta = 1;
				END IF;
				IF iRespuesta = 1 THEN
					--cSql:=cSql || 'UPDATE ctl_polizas SET fecha = NOW()::date, sku =  '|| iSku ||', cantidad = '|| iCantidad  ||'';
					UPDATE ctl_polizas SET fecha = NOW()::date,sku = iSku,cantidad = iCantidad WHERE id_polizas = iPoliza;
				END IF;
			ELSE
				iRespuesta = 0;
			END IF;
			
	ELSE
		iRespuesta = 0;
	END IF;
	
	IF EXISTS (SELECT FROM ctl_empleado WHERE id_empleado = iEmpleado) THEN --si existe el empleado
		UPDATE ctl_empleado SET nombre = sNombre,apellido = sApellido,puesto = sPuesto WHERE id_empleado = iEmpleado;
		iRespuesta = 1;
	ELSE
		iRespuesta= 0;
	END IF;

	RETURN iRespuesta;
	--EXECUTE cSql;
END;
$BODY$
LANGUAGE plpgsql VOLATILE SECURITY DEFINER;

















ricardo
https://cronohub.com/ficha/7dc550c83c75a49032148a7f90f248cf

marisela
https://cronohub.com/ficha/2e17b55bd1f302b773d04e4938adef85

lluvia
https://cronohub.com/ficha/95f2351eb791917e3e1cee823cf25969

?token=c8f1e1d2632f3a5af838f86190eb7695&tienda=3&caja=30&empleado=98493701&ip=http://npvocdev43.coppel.io:20541/&ipcaja=10.59.21.109&sistema=1&impresora=BIXOLON SAMSUNG SRP-350plus&nombreempleado=MARIO PRUEA PRUEBA&cliente=98493701&iptienda=10.40.116.128&foliopreventa=229472&fechatienda=2021-09-30&flagVentaContado=1
?token=c8f1e1d2632f3a5af838f86190eb7695&tienda=3&caja=30&empleado=98493701&ip=http://npvocdev43.coppel.io:20541/&ipcaja=10.59.21.109&sistema=1&impresora=BIXOLON SAMSUNG SRP-350plus&nombreempleado=MARIO PRUEA PRUEBA&cliente=98493701&iptienda=10.40.116.128&foliopreventa=229472&fechatienda=2021-09-30&flagVentaContado=1





