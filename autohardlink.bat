@ECHO ON

SET SrcRoot=C:\Users\Lili\Documents\GitHub\YAGPDB-cc-rp.wiki
SET TargetRoot=C:\Users\Lili\Documents\GitHub\Projet-Nucleus-Custom-Command.wiki

FOR /D %%A IN ("%SrcRoot%\*") DO (
    MKLINK /D "%TargetRoot%\%%~NA" "%%~A"
    )

FOR %%A IN ("%SrcRoot%\*") DO (
    MKLINK /H "%TargetRoot%\%%~NXA" "%%~A"
    )

PAUSE
EXIT