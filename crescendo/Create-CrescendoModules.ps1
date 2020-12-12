$files = Get-ChildItem -Path ./json -Filter "*.json"
$null = Get-ChildItem -Path ./psm1 -Filter "*.psm1" | Remove-Item
if (Test-Path GSM.psm1) {
    $null = Remove-Item GSM.psm1
}
foreach($file in $files) {
    Export-CrescendoModule -ConfigurationFile $file.FullName -ModuleName ("./psm1/" + ($file.Name).Replace(".json",".psm1"))
}
$modules = Get-ChildItem -Path ./psm1 -Filter "*.psm1"
foreach($module in $modules) {
    $moduleName = ($module.Name).Replace(".psm1","")
    Remove-Module $moduleName -ErrorAction Ignore -WarningAction Ignore
    Import-Module $module.FullName -Global -WarningAction Ignore
    $foo = ($moduleName).Split("_")
    $fName = $foo[1] + "-" + $foo[0]
    $command = Get-Command $fName
    "Function " + $command + " {`n`n" + $command.Definition  + "}`n`n" >> GSM.psm1
}
