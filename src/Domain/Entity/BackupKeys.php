<?php

declare(strict_types=1);

namespace Backup8\Domain\Entity;

use JsonSerializable;
use Backup8\Domain\ValueObject\Iam\AccessKeyId;
use Backup8\Domain\ValueObject\Iam\SecretAccessKey;

/**
 * @OA\Schema(
 *   title="BackupKeys",
 *   required={"accessKeyId", "secretAccessKey"}
 * )
 */
class BackupKeys implements JsonSerializable
{
    /**
     * @OA\Property(type="string")
     */
    private AccessKeyId $accessKeyId;

    /**
     * @OA\Property(type="string")
     */
    private SecretAccessKey $secretAccessKey;

    public function __construct(
        AccessKeyId $accessKeyId,
        SecretAccessKey $secretAccessKey
    ) {
        $this->accessKeyId = $accessKeyId;
        $this->secretAccessKey = $secretAccessKey;
    }

    public function getAccessKeyId(): AccessKeyId
    {
        return $this->accessKeyId;
    }

    public function getSecretAccessKey(): SecretAccessKey
    {
        return $this->secretAccessKey;
    }

    public function jsonSerialize(): array
    {
        return [
            "accessKeyId" => $this->accessKeyId,
            "secretAccessKey" => $this->secretAccessKey,
        ];
    }
}
