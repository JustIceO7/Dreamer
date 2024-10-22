package sd

import (
	"encoding/json"
	"fmt"

	"github.com/Strum355/log"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type jsonTextToImageResponse struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}

type jsonInfoResponse struct {
	AllPrompts         []string `json:"all_prompts"`
	AllNegativePrompts []string `json:"all_negative_prompts"`
	AllSeeds           []int    `json:"all_seeds"`
	SDModelName        string   `json:"sd_model_name"`
	Width              int      `json:"width"`
	Height             int      `json:"height"`
}

type TextToImageResponse struct {
	Images []string `json:"images"`
	Seeds  []int    `json:"seeds"`
}

type Images struct {
	Image [][]byte
}

func (img *Images) AppendImage(newData []byte) {
	img.Image = append(img.Image, newData)
}

type TextToImageRequest struct {
	Prompt                            string                 `json:"prompt"`
	NegativePrompt                    string                 `json:"negative_prompt,omitempty"`
	Styles                            []string               `json:"styles,omitempty"`
	Seed                              int                    `json:"seed,omitempty"`
	Subseed                           int                    `json:"subseed,omitempty"`
	SubseedStrength                   float64                `json:"subseed_strength,omitempty"`
	SeedResizeFromH                   int                    `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   int                    `json:"seed_resize_from_w,omitempty"`
	SamplerName                       string                 `json:"sampler_name,omitempty"`
	Scheduler                         string                 `json:"scheduler,omitempty"`
	BatchSize                         int                    `json:"batch_size,omitempty"`
	NIter                             int                    `json:"n_iter,omitempty"`
	Steps                             int                    `json:"steps,omitempty"`
	CfgScale                          float64                `json:"cfg_scale,omitempty"`
	Width                             int                    `json:"width,omitempty"`
	Height                            int                    `json:"height,omitempty"`
	RestoreFaces                      bool                   `json:"restore_faces,omitempty"`
	Tiling                            bool                   `json:"tiling,omitempty"`
	DoNotSaveSamples                  bool                   `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid                     bool                   `json:"do_not_save_grid,omitempty"`
	Eta                               float64                `json:"eta,omitempty"`
	DenoisingStrength                 float64                `json:"denoising_strength,omitempty"`
	SMinUncond                        float64                `json:"s_min_uncond,omitempty"`
	SChurn                            float64                `json:"s_churn,omitempty"`
	STmax                             float64                `json:"s_tmax,omitempty"`
	STmin                             float64                `json:"s_tmin,omitempty"`
	SNoise                            float64                `json:"s_noise,omitempty"`
	OverrideSettings                  map[string]interface{} `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards bool                   `json:"override_settings_restore_afterwards,omitempty"`
	RefinerCheckpoint                 string                 `json:"refiner_checkpoint,omitempty"`
	RefinerSwitchAt                   float64                `json:"refiner_switch_at,omitempty"`
	DisableExtraNetworks              bool                   `json:"disable_extra_networks,omitempty"`
	FirstpassImage                    string                 `json:"firstpass_image,omitempty"`
	Comments                          map[string]interface{} `json:"comments,omitempty"`
	EnableHr                          bool                   `json:"enable_hr,omitempty"`
	FirstphaseWidth                   int                    `json:"firstphase_width,omitempty"`
	FirstphaseHeight                  int                    `json:"firstphase_height,omitempty"`
	HrScale                           float64                `json:"hr_scale,omitempty"`
	HrUpscaler                        string                 `json:"hr_upscaler,omitempty"`
	HrSecondPassSteps                 int                    `json:"hr_second_pass_steps,omitempty"`
	HrResizeX                         int                    `json:"hr_resize_x,omitempty"`
	HrResizeY                         int                    `json:"hr_resize_y,omitempty"`
	HrCheckpointName                  string                 `json:"hr_checkpoint_name,omitempty"`
	HrSamplerName                     string                 `json:"hr_sampler_name,omitempty"`
	HrScheduler                       string                 `json:"hr_scheduler,omitempty"`
	HrPrompt                          string                 `json:"hr_prompt,omitempty"`
	HrNegativePrompt                  string                 `json:"hr_negative_prompt,omitempty"`
	ForceTaskID                       string                 `json:"force_task_id,omitempty"`
	SamplerIndex                      string                 `json:"sampler_index,omitempty"`
	ScriptName                        string                 `json:"script_name,omitempty"`
	ScriptArgs                        []interface{}          `json:"script_args,omitempty"`
	SendImages                        bool                   `json:"send_images,omitempty"`
	SaveImages                        bool                   `json:"save_images,omitempty"`
	AlwaysonScripts                   map[string]interface{} `json:"alwayson_scripts,omitempty"`
	Infotext                          string                 `json:"infotext,omitempty"`
}

// Handles text2image requests
func TextToImage(prompt string) (*TextToImageResponse, error) {
	client := resty.New()

	requestBody := TextToImageRequest{
		Prompt:         prompt,
		NegativePrompt: "",
		Width:          1024,
		Height:         1024,
		Steps:          25,
		CfgScale:       7.5,
		SamplerName:    "Euler a",
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(viper.GetString("sd.url") + "sdapi/v1/txt2img")

	if err != nil {
		log.WithError(err).Error("Failed to send text2image request.")
		return nil, err
	}
	fmt.Println("Response Status Code:", resp.StatusCode())

	// Unmarshal response body
	generateResponse := &jsonTextToImageResponse{}
	if err := json.Unmarshal(resp.Body(), &generateResponse); err != nil {
		log.WithError(err).Error("Failed to unmarshal text2image body response.")
		return nil, err
	}

	infoStruct := &jsonInfoResponse{}
	err = json.Unmarshal([]byte(generateResponse.Info), infoStruct)
	if err != nil {
		log.WithError(err).Error("Failed to unmarshal text2image info response")
		return nil, err
	}

	return &TextToImageResponse{
		Images: generateResponse.Images,
		Seeds:  infoStruct.AllSeeds}, nil
}
