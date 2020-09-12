@ECHO ON

SET SrcRoot=C:\Users\Lili\Documents\GitHub\YAGPDB-cc-rp\documentations
SET TargetRoot=C:\Users\Lili\Documents\GitHub\Projet-Nucleus-Custom-Command.wiki
SET RepoRoot=C:\Users\Lili\Documents\GitHub\YAGPDB-cc-rp.wiki

FOR /D %%A IN ("%SrcRoot%\*") DO (
    MKLINK /D "%TargetRoot%\%%~NA" "%%~A"
	MKLINK /D "%RepoRoot%\%%~NA" "%%~A"
    )

FOR %%A IN ("%SrcRoot%\*") DO (
    MKLINK /H "%TargetRoot%\%%~NXA" "%%~A"
MKLINK /H "%RepoRoot%\%%~NXA" "%%~A"
    )

PAUSE
EXIT