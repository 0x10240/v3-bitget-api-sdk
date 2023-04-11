<?php
/**
 * MarginCrossAccountApiTest
 * PHP version 7.4
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Bitget Open API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 6.2.1
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Please update the test case below to test the endpoint.
 */

namespace Bitget\Test\Api;

use \Bitget\Configuration;
use \Bitget\ApiException;
use \Bitget\ObjectSerializer;
use Exception;
use PHPUnit\Framework\TestCase;

/**
 * MarginCrossAccountApiTest Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class MarginCrossAccountApiTest extends TestCase
{

    private $config;

    private $apiInstance;

    /**
     * Setup before running any test cases
     */
    public static function setUpBeforeClass(): void
    {
    }

    /**
     * Setup before running each test case
     */
    public function setUp(): void
    {
        $this->config = \Bitget\Config::getDefaultConfig();
        $this->apiInstance = new \Bitget\Api\MarginCrossAccountApi(
        // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
        // This is optional, `GuzzleHttp\Client` will be used as default.
            new \GuzzleHttp\Client(),
            $this->config
        );
    }

    /**
     * Clean up after running each test case
     */
    public function tearDown(): void
    {
    }

    /**
     * Clean up after running all test cases
     */
    public static function tearDownAfterClass(): void
    {
    }

    /**
     * Test case for marginCrossAccountAssets
     *
     * assets.
     *
     */
    public function testMarginCrossAccountAssets()
    {
        try {
            $result = $this->apiInstance->marginCrossAccountAssets("USDT");
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            foreach ($result->getData() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertEquals($item->getCoin(), "USDT");
                $this->assertNotNull($item->getAvailable());
                $this->assertNotNull($item->getBorrow());
                $this->assertNotNull($item->getFrozen());
                $this->assertNotNull($item->getInterest());
                $this->assertNotNull($item->getModelName());
                $this->assertNotNull($item->getNet());
                $this->assertNotNull($item->getTotalAmount());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossAccountBorrow
     *
     * borrow.
     *
     */
    public function testMarginCrossAccountBorrow()
    {
        try {
            $req = new \Bitget\Model\MarginCrossLimitRequest(); //
            $req->setCoin("USDT");
            $req->setBorrowAmount("1");
            $result = $this->apiInstance->marginCrossAccountBorrow($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertEquals($result->getData()->getCoin(), "USDT");
            $this->assertNotNull($result->getData()->getBorrowAmount());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossAccountMaxBorrowableAmount
     *
     * maxBorrowableAmount.
     *
     */
    public function testMarginCrossAccountMaxBorrowableAmount()
    {
        try {
            $req = new \Bitget\Model\MarginCrossMaxBorrowRequest(); //
            $req->setCoin("USDT");
            $result = $this->apiInstance->marginCrossAccountMaxBorrowableAmount($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertEquals($result->getData()->getCoin(), "USDT");
            $this->assertNotNull($result->getData()->getMaxBorrowableAmount());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossAccountMaxTransferOutAmount
     *
     * maxTransferOutAmount.
     *
     */
    public function testMarginCrossAccountMaxTransferOutAmount()
    {
        try {
            $result = $this->apiInstance->marginCrossAccountMaxTransferOutAmount("USDT");
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertEquals($result->getData()->getCoin(), "USDT");
            $this->assertNotNull($result->getData()->getMaxTransferOutAmount());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossAccountRepay
     *
     * repay.
     *
     */
    public function testMarginCrossAccountRepay()
    {
        try {
            $req = new \Bitget\Model\MarginCrossRepayRequest(); //
            $req->setCoin("USDT");
            $req->setRepayAmount("1");
            $result = $this->apiInstance->marginCrossAccountRepay($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertEquals($result->getData()->getCoin(), "USDT");
            $this->assertNotNull($result->getData()->getRepayAmount());
            $this->assertNotNull($result->getData()->getRemainDebtAmount());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossAccountRiskRate
     *
     * riskRate.
     *
     */
    public function testMarginCrossAccountRiskRate()
    {
        try {
            $result = $this->apiInstance->marginCrossAccountRiskRate();
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getRiskRate());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }
}