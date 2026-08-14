# wintree shell integration
function wintree-cd {

        $result = & "_winTree_" cd @args

        if ($result -and (Test-Path $result -PathType Container)) {
            Set-Location $result
        }

       
}
