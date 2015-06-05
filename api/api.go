package api

import (
    "log"
    "net/http"
    "net/url"
    "crypto/tls"
    "io/ioutil"
    "strings"
    "os"
    "encoding/json"
    "path/filepath"
        
    "github.com/krustnic/runtime-info/info"
)

var client *http.Client

type LocalConfig struct {
    Username string
    Password string
    ApiHost  string
    ConfigApiPath string
}

var Config *LocalConfig
var localConfigPath = getExecutablePath() + "/config.json"

func init() {
    transport := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify : true},
    }
    
    client = &http.Client{ Transport : transport }

    // Check for manually specified config file
    if len( os.Args ) == 3 && os.Args[1] == "-c" {
        localConfigPath = os.Args[2]
    }
    log.Printf("Local config file path: %s", localConfigPath )
    
    // Load local config
    config, err := LoadLocalConfig()
    if err != nil {
        os.Exit(1)
    }
    
    Config = &config   
}

func LoadLocalConfig() (LocalConfig, error) {
    var localConfig = LocalConfig{}
    
    configPath := localConfigPath
    
    log.Printf("Start loading local config file: %s", configPath)
    
    contentBytes, err := ioutil.ReadFile( configPath )
        
    if err != nil {
        log.Printf( "Error reading local settings file: %s", err )
        return localConfig, err
    }
    
    if err := json.Unmarshal(contentBytes, &localConfig); err != nil {
        log.Printf( "Error parsing local settings file from JSON to struct: %s", err )
        return localConfig, err
    }    
    
    log.Printf("Loading config file is complite: %+v", localConfig)        
    
    return localConfig, err
}

func getExecutablePath() string {
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))        
    return dir
}

func RequestRemoteConfiguration() string {    
    data := map[string]string{ "data" : info.ToJSON() }
    
    content, _ := Request( "POST", Config.ConfigApiPath, data )    
    return content    
}

func Request( method, path string, data map[string]string ) (string, error) {
    form := url.Values{}
    
    for key, value := range data {
        form.Add(key, value)    
    }
    
    apiUrl := Config.ApiHost + path
    
    log.Printf("Create API request: %s, %s", method, apiUrl)
    
    req, err := http.NewRequest(method, apiUrl, strings.NewReader(form.Encode()))
    req.SetBasicAuth( Config.Username, Config.Password )
    
    if err != nil {
        log.Printf("Error on api request: %s", err)   
        return "", err
    }
    
    res, err := client.Do(req)
    
    if err != nil {
        log.Printf("Error on response: %s", err)
        return "", err
    }
    
    if ( res.StatusCode != 200 ) {
        log.Printf("Error. Not success response status code: %d", res.StatusCode)
        return "", err
    }
    
    defer res.Body.Close()
    contents, err := ioutil.ReadAll(res.Body)
    
    bodyResponse := string( contents )
    
    log.Printf("Got API response: %s", bodyResponse)
        
    return bodyResponse, err
}
