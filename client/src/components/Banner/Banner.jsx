import bannerImage from '../../assets/images/zalora_banner.png';

export const Banner = () => {
    return (
        <div className="banner">
            <img src={bannerImage} alt="banner" style={{ width: "100%" }}/>
        </div>
    );
}